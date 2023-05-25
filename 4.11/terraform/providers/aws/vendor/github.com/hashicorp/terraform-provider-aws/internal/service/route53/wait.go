package route53

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	changeTimeout      = 30 * time.Minute
	changeMinTimeout   = 5 * time.Second
	changePollInterval = 15 * time.Second
	changeMinDelay     = 10
	changeMaxDelay     = 30

	hostedZoneDNSSECStatusTimeout = 5 * time.Minute

	keySigningKeyStatusTimeout = 5 * time.Minute
)

func waitChangeInfoStatusInsync(conn *route53.Route53, changeID string) (*route53.ChangeInfo, error) { //nolint:unparam
	rand.Seed(time.Now().UTC().UnixNano())

	// Route53 is vulnerable to throttling so longer delays, poll intervals helps significantly to avoid

	stateConf := &resource.StateChangeConf{
		Pending:      []string{route53.ChangeStatusPending},
		Target:       []string{route53.ChangeStatusInsync},
		Delay:        time.Duration(rand.Int63n(changeMaxDelay-changeMinDelay)+changeMinDelay) * time.Second,
		MinTimeout:   changeMinTimeout,
		PollInterval: changePollInterval,
		Refresh:      statusChangeInfo(conn, changeID),
		Timeout:      changeTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*route53.ChangeInfo); ok {
		return output, err
	}

	return nil, err
}

func waitHostedZoneDNSSECStatusUpdated(conn *route53.Route53, hostedZoneID string, status string) (*route53.DNSSECStatus, error) { //nolint:unparam
	stateConf := &resource.StateChangeConf{
		Target:     []string{status},
		Refresh:    statusHostedZoneDNSSEC(conn, hostedZoneID),
		MinTimeout: 5 * time.Second,
		Timeout:    hostedZoneDNSSECStatusTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*route53.DNSSECStatus); ok {
		if err != nil && output != nil && output.ServeSignature != nil && output.StatusMessage != nil {
			newErr := fmt.Errorf("%s: %s", aws.StringValue(output.ServeSignature), aws.StringValue(output.StatusMessage))

			switch e := err.(type) {
			case *resource.TimeoutError:
				if e.LastError == nil {
					e.LastError = newErr
				}
			case *resource.UnexpectedStateError:
				if e.LastError == nil {
					e.LastError = newErr
				}
			}
		}

		return output, err
	}

	return nil, err
}

func waitKeySigningKeyStatusUpdated(conn *route53.Route53, hostedZoneID string, name string, status string) (*route53.KeySigningKey, error) { //nolint:unparam
	stateConf := &resource.StateChangeConf{
		Target:     []string{status},
		Refresh:    statusKeySigningKey(conn, hostedZoneID, name),
		MinTimeout: 5 * time.Second,
		Timeout:    keySigningKeyStatusTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*route53.KeySigningKey); ok {
		if err != nil && output != nil && output.Status != nil && output.StatusMessage != nil {
			newErr := fmt.Errorf("%s: %s", aws.StringValue(output.Status), aws.StringValue(output.StatusMessage))

			var te *resource.TimeoutError
			var use *resource.UnexpectedStateError
			if ok := errors.As(err, &te); ok && te.LastError == nil {
				te.LastError = newErr
			} else if ok := errors.As(err, &use); ok && use.LastError == nil {
				use.LastError = newErr
			}
		}

		return output, err
	}

	return nil, err
}
