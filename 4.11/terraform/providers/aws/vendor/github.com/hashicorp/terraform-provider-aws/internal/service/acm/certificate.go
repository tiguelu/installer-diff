package acm

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

const (
	// Maximum amount of time for ACM Certificate cross-service reference propagation.
	// Removal of ACM Certificates from API Gateway Custom Domains can take >15 minutes.
	AcmCertificateCrossServicePropagationTimeout = 20 * time.Minute

	// Maximum amount of time for ACM Certificate asynchronous DNS validation record assignment.
	// This timeout is unrelated to any creation or validation of those assigned DNS records.
	AcmCertificateDnsValidationAssignmentTimeout = 5 * time.Minute
)

func ResourceCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateCreate,
		Read:   resourceCertificateRead,
		Update: resourceCertificateUpdate,
		Delete: resourceCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"certificate_body": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_chain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"certificate_authority_arn": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"domain_name": {
				// AWS Provider 3.0.0 aws_route53_zone references no longer contain a
				// trailing period, no longer requiring a custom StateFunc
				// to prevent ACM API error
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"private_key", "certificate_body", "certificate_chain"},
				ValidateFunc:  validation.StringDoesNotMatch(regexp.MustCompile(`\.$`), "cannot end with a period"),
			},
			"subject_alternative_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					// AWS Provider 3.0.0 aws_route53_zone references no longer contain a
					// trailing period, no longer requiring a custom StateFunc
					// to prevent ACM API error
					Type: schema.TypeString,
					ValidateFunc: validation.All(
						validation.StringLenBetween(1, 253),
						validation.StringDoesNotMatch(regexp.MustCompile(`\.$`), "cannot end with a period"),
					),
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"private_key", "certificate_body", "certificate_chain"},
			},
			"validation_method": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"private_key", "certificate_body", "certificate_chain", "certificate_authority_arn"},
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_validation_options": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_record_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_record_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Set: acmDomainValidationOptionsHash,
			},
			"validation_emails": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"options": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if _, ok := d.GetOk("private_key"); ok {
						// ignore diffs for imported certs; they have a different logging preference
						// default to requested certs which can't be changed by the ImportCertificate API
						return true
					}
					// behave just like verify.SuppressMissingOptionalConfigurationBlock() for requested certs
					return old == "1" && new == "0"
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_transparency_logging_preference": {
							Type:          schema.TypeString,
							Optional:      true,
							Default:       acm.CertificateTransparencyLoggingPreferenceEnabled,
							ForceNew:      true,
							ConflictsWith: []string{"private_key", "certificate_body", "certificate_chain"},
							ValidateFunc: validation.StringInSlice([]string{
								acm.CertificateTransparencyLoggingPreferenceEnabled,
								acm.CertificateTransparencyLoggingPreferenceDisabled,
							}, false),
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				// Attempt to calculate the domain validation options based on domains present in domain_name and subject_alternative_names
				if diff.Get("validation_method").(string) == "DNS" && (diff.HasChange("domain_name") || diff.HasChange("subject_alternative_names")) {
					domainValidationOptionsList := []interface{}{map[string]interface{}{
						// AWS Provider 3.0 -- plan-time validation prevents "domain_name"
						// argument to accept a string with trailing period; thus, trim of trailing period
						// no longer required here
						"domain_name": diff.Get("domain_name").(string),
					}}

					if sanSet, ok := diff.Get("subject_alternative_names").(*schema.Set); ok {
						for _, sanRaw := range sanSet.List() {
							san, ok := sanRaw.(string)

							if !ok {
								continue
							}

							m := map[string]interface{}{
								// AWS Provider 3.0 -- plan-time validation prevents "subject_alternative_names"
								// argument to accept strings with trailing period; thus, trim of trailing period
								// no longer required here
								"domain_name": san,
							}

							domainValidationOptionsList = append(domainValidationOptionsList, m)
						}
					}

					if err := diff.SetNew("domain_validation_options", schema.NewSet(acmDomainValidationOptionsHash, domainValidationOptionsList)); err != nil {
						return fmt.Errorf("error setting new domain_validation_options diff: %w", err)
					}
				}

				return nil
			},
			verify.SetTagsDiff,
		),
	}
}

func resourceCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	if _, ok := d.GetOk("domain_name"); ok {
		if _, ok := d.GetOk("certificate_authority_arn"); ok {
			return resourceCertificateCreateRequested(d, meta)
		}

		if _, ok := d.GetOk("validation_method"); !ok {
			return errors.New("validation_method must be set when creating a certificate")
		}
		return resourceCertificateCreateRequested(d, meta)
	} else if _, ok := d.GetOk("private_key"); ok {
		if _, ok := d.GetOk("certificate_body"); !ok {
			return errors.New("certificate_body must be set when importing a certificate with private_key")
		}
		return resourceCertificateCreateImported(d, meta)
	}
	return errors.New("certificate must be imported (private_key) or created (domain_name)")
}

func resourceCertificateCreateImported(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ACMConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(d.Get("tags").(map[string]interface{})))

	input := &acm.ImportCertificateInput{
		Certificate: []byte(d.Get("certificate_body").(string)),
		PrivateKey:  []byte(d.Get("private_key").(string)),
	}

	if v, ok := d.GetOk("certificate_chain"); ok {
		input.CertificateChain = []byte(v.(string))
	}

	if len(tags) > 0 {
		input.Tags = Tags(tags.IgnoreAWS())
	}

	output, err := conn.ImportCertificate(input)

	if err != nil {
		return fmt.Errorf("error importing ACM Certificate: %w", err)
	}

	d.SetId(aws.StringValue(output.CertificateArn))

	return resourceCertificateRead(d, meta)
}

func resourceCertificateCreateRequested(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ACMConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(d.Get("tags").(map[string]interface{})))

	params := &acm.RequestCertificateInput{
		DomainName:       aws.String(d.Get("domain_name").(string)),
		IdempotencyToken: aws.String(resource.PrefixedUniqueId("tf")), // 32 character limit
		Options:          expandAcmCertificateOptions(d.Get("options").([]interface{})),
	}

	if len(tags) > 0 {
		params.Tags = Tags(tags.IgnoreAWS())
	}

	if caARN, ok := d.GetOk("certificate_authority_arn"); ok {
		params.CertificateAuthorityArn = aws.String(caARN.(string))
	}

	if sans, ok := d.GetOk("subject_alternative_names"); ok {
		subjectAlternativeNames := make([]*string, len(sans.(*schema.Set).List()))
		for i, sanRaw := range sans.(*schema.Set).List() {
			subjectAlternativeNames[i] = aws.String(sanRaw.(string))
		}
		params.SubjectAlternativeNames = subjectAlternativeNames
	}

	if v, ok := d.GetOk("validation_method"); ok {
		params.ValidationMethod = aws.String(v.(string))
	}

	log.Printf("[DEBUG] ACM Certificate Request: %#v", params)
	resp, err := conn.RequestCertificate(params)

	if err != nil {
		return fmt.Errorf("Error requesting certificate: %s", err)
	}

	d.SetId(aws.StringValue(resp.CertificateArn))

	return resourceCertificateRead(d, meta)
}

func resourceCertificateRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ACMConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	params := &acm.DescribeCertificateInput{
		CertificateArn: aws.String(d.Id()),
	}

	return resource.Retry(AcmCertificateDnsValidationAssignmentTimeout, func() *resource.RetryError {
		resp, err := conn.DescribeCertificate(params)

		if !d.IsNewResource() && tfawserr.ErrCodeEquals(err, acm.ErrCodeResourceNotFoundException) {
			log.Printf("[WARN] ACM Certificate (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error reading ACM Certificate (%s): %w", d.Id(), err))
		}

		if resp == nil || resp.Certificate == nil {
			return resource.NonRetryableError(fmt.Errorf("error reading ACM Certificate (%s): empty response", d.Id()))
		}

		if !d.IsNewResource() && aws.StringValue(resp.Certificate.Status) == acm.CertificateStatusValidationTimedOut {
			log.Printf("[WARN] ACM Certificate (%s) validation timed out, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		d.Set("domain_name", resp.Certificate.DomainName)
		d.Set("arn", resp.Certificate.CertificateArn)
		d.Set("certificate_authority_arn", resp.Certificate.CertificateAuthorityArn)

		if err := d.Set("subject_alternative_names", cleanUpSubjectAlternativeNames(resp.Certificate)); err != nil {
			return resource.NonRetryableError(err)
		}

		domainValidationOptions, emailValidationOptions, err := convertValidationOptions(resp.Certificate)

		if err != nil {
			return resource.RetryableError(err)
		}

		if err := d.Set("domain_validation_options", domainValidationOptions); err != nil {
			return resource.NonRetryableError(err)
		}
		if err := d.Set("validation_emails", emailValidationOptions); err != nil {
			return resource.NonRetryableError(err)
		}

		d.Set("validation_method", resourceCertificateValidationMethod(resp.Certificate))

		if err := d.Set("options", flattenAcmCertificateOptions(resp.Certificate.Options)); err != nil {
			return resource.NonRetryableError(fmt.Errorf("error setting certificate options: %s", err))
		}

		d.Set("status", resp.Certificate.Status)

		tags, err := ListTags(conn, d.Id())

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error listing tags for ACM Certificate (%s): %s", d.Id(), err))
		}

		tags = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig)

		//lintignore:AWSR002
		if err := d.Set("tags", tags.RemoveDefaultConfig(defaultTagsConfig).Map()); err != nil {
			return resource.NonRetryableError(fmt.Errorf("error setting tags: %w", err))
		}

		if err := d.Set("tags_all", tags.Map()); err != nil {
			return resource.NonRetryableError(fmt.Errorf("error setting tags_all: %w", err))
		}

		return nil
	})
}

func resourceCertificateValidationMethod(certificate *acm.CertificateDetail) string {
	if aws.StringValue(certificate.Type) == acm.CertificateTypeAmazonIssued {
		for _, domainValidation := range certificate.DomainValidationOptions {
			if domainValidation.ValidationMethod != nil {
				return aws.StringValue(domainValidation.ValidationMethod)
			}
		}
	}

	return "NONE"
}

func resourceCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ACMConn

	if d.HasChanges("private_key", "certificate_body", "certificate_chain") {
		// Prior to version 3.0.0 of the Terraform AWS Provider, these attributes were stored in state as hashes.
		// If the changes to these attributes are only changes only match updating the state value, then skip the API call.
		oCBRaw, nCBRaw := d.GetChange("certificate_body")
		oCCRaw, nCCRaw := d.GetChange("certificate_chain")
		oPKRaw, nPKRaw := d.GetChange("private_key")

		if !isChangeNormalizeCertRemoval(oCBRaw, nCBRaw) || !isChangeNormalizeCertRemoval(oCCRaw, nCCRaw) || !isChangeNormalizeCertRemoval(oPKRaw, nPKRaw) {
			input := &acm.ImportCertificateInput{
				Certificate:    []byte(d.Get("certificate_body").(string)),
				CertificateArn: aws.String(d.Get("arn").(string)),
				PrivateKey:     []byte(d.Get("private_key").(string)),
			}

			if chain, ok := d.GetOk("certificate_chain"); ok {
				input.CertificateChain = []byte(chain.(string))
			}

			_, err := conn.ImportCertificate(input)

			if err != nil {
				return fmt.Errorf("error re-importing ACM Certificate (%s): %w", d.Id(), err)
			}
		}
	}

	if d.HasChange("tags_all") {
		o, n := d.GetChange("tags_all")
		if err := UpdateTags(conn, d.Id(), o, n); err != nil {
			return fmt.Errorf("error updating tags: %s", err)
		}
	}
	return resourceCertificateRead(d, meta)
}

func cleanUpSubjectAlternativeNames(cert *acm.CertificateDetail) []string {
	sans := cert.SubjectAlternativeNames
	vs := make([]string, 0)
	for _, v := range sans {
		if aws.StringValue(v) != aws.StringValue(cert.DomainName) {
			vs = append(vs, aws.StringValue(v))
		}
	}
	return vs
}

func convertValidationOptions(certificate *acm.CertificateDetail) ([]map[string]interface{}, []string, error) {
	var domainValidationResult []map[string]interface{}
	var emailValidationResult []string

	switch aws.StringValue(certificate.Type) {
	case acm.CertificateTypeAmazonIssued:
		if len(certificate.DomainValidationOptions) == 0 && aws.StringValue(certificate.Status) == acm.DomainStatusPendingValidation {
			log.Printf("[DEBUG] No validation options need to retry.")
			return nil, nil, fmt.Errorf("No validation options need to retry.")
		}
		for _, o := range certificate.DomainValidationOptions {
			if o.ResourceRecord != nil {
				validationOption := map[string]interface{}{
					"domain_name":           aws.StringValue(o.DomainName),
					"resource_record_name":  aws.StringValue(o.ResourceRecord.Name),
					"resource_record_type":  aws.StringValue(o.ResourceRecord.Type),
					"resource_record_value": aws.StringValue(o.ResourceRecord.Value),
				}
				domainValidationResult = append(domainValidationResult, validationOption)
			} else if o.ValidationEmails != nil && len(o.ValidationEmails) > 0 {
				for _, validationEmail := range o.ValidationEmails {
					emailValidationResult = append(emailValidationResult, *validationEmail)
				}
			} else if o.ValidationStatus == nil || aws.StringValue(o.ValidationStatus) == acm.DomainStatusPendingValidation {
				log.Printf("[DEBUG] Asynchronous ACM service domain validation assignment not complete, need to retry: %#v", o)
				return nil, nil, fmt.Errorf("asynchronous ACM service domain validation assignment not complete, need to retry: %#v", o)
			}
		}
	case acm.CertificateTypePrivate:
		// While ACM PRIVATE certificates do not need to be validated, there is a slight delay for
		// the API to fill in all certificate details, which is during the PENDING_VALIDATION status.
		if aws.StringValue(certificate.Status) == acm.DomainStatusPendingValidation {
			return nil, nil, fmt.Errorf("certificate still pending issuance")
		}
	}

	return domainValidationResult, emailValidationResult, nil
}

func resourceCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ACMConn

	log.Printf("[INFO] Deleting ACM Certificate: %s", d.Id())

	params := &acm.DeleteCertificateInput{
		CertificateArn: aws.String(d.Id()),
	}

	err := resource.Retry(AcmCertificateCrossServicePropagationTimeout, func() *resource.RetryError {
		_, err := conn.DeleteCertificate(params)

		if tfawserr.ErrCodeEquals(err, acm.ErrCodeResourceInUseException) {
			return resource.RetryableError(err)
		}

		if err != nil {
			return resource.NonRetryableError(err)
		}

		return nil
	})

	if tfresource.TimedOut(err) {
		_, err = conn.DeleteCertificate(params)
	}

	if tfawserr.ErrCodeEquals(err, acm.ErrCodeResourceNotFoundException) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("error deleting ACM Certificate (%s): %w", d.Id(), err)
	}

	return nil
}

func acmDomainValidationOptionsHash(v interface{}) int {
	m, ok := v.(map[string]interface{})

	if !ok {
		return 0
	}

	if v, ok := m["domain_name"].(string); ok {
		return create.StringHashcode(v)
	}

	return 0
}

func expandAcmCertificateOptions(l []interface{}) *acm.CertificateOptions {
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	m := l[0].(map[string]interface{})

	options := &acm.CertificateOptions{}

	if v, ok := m["certificate_transparency_logging_preference"]; ok {
		options.CertificateTransparencyLoggingPreference = aws.String(v.(string))
	}

	return options
}

func flattenAcmCertificateOptions(co *acm.CertificateOptions) []interface{} {
	m := map[string]interface{}{
		"certificate_transparency_logging_preference": aws.StringValue(co.CertificateTransparencyLoggingPreference),
	}

	return []interface{}{m}
}

func isChangeNormalizeCertRemoval(oldRaw, newRaw interface{}) bool {
	old, ok := oldRaw.(string)

	if !ok {
		return false
	}

	new, ok := newRaw.(string)

	if !ok {
		return false
	}

	newCleanVal := sha1.Sum(stripCR([]byte(strings.TrimSpace(new))))
	return hex.EncodeToString(newCleanVal[:]) == old
}

// strip CRs from raw literals. Lifted from go/scanner/scanner.go
// See https://github.com/golang/go/blob/release-branch.go1.6/src/go/scanner/scanner.go#L479
func stripCR(b []byte) []byte {
	c := make([]byte, len(b))
	i := 0
	for _, ch := range b {
		if ch != '\r' {
			c[i] = ch
			i++
		}
	}
	return c[:i]
}
