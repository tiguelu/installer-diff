// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Crls in the authenticated account and Amazon Web Services Region.
// Required permissions: rolesanywhere:ListCrls.
func (c *Client) ListCrls(ctx context.Context, params *ListCrlsInput, optFns ...func(*Options)) (*ListCrlsOutput, error) {
	if params == nil {
		params = &ListCrlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCrls", params, optFns, c.addOperationListCrlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCrlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCrlsInput struct {

	// A token that indicates where the output should continue from, if a previous
	// operation did not show all results. To get the next results, call the operation
	// again with this value.
	NextToken *string

	// The number of resources in the paginated list.
	PageSize *int32

	noSmithyDocumentSerde
}

type ListCrlsOutput struct {

	// A list of certificate revocation lists (CRL).
	Crls []types.CrlDetail

	// A token that indicates where the output should continue from, if a previous
	// operation did not show all results. To get the next results, call the operation
	// again with this value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCrlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCrls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCrls{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCrls(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListCrlsAPIClient is a client that implements the ListCrls operation.
type ListCrlsAPIClient interface {
	ListCrls(context.Context, *ListCrlsInput, ...func(*Options)) (*ListCrlsOutput, error)
}

var _ ListCrlsAPIClient = (*Client)(nil)

// ListCrlsPaginatorOptions is the paginator options for ListCrls
type ListCrlsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCrlsPaginator is a paginator for ListCrls
type ListCrlsPaginator struct {
	options   ListCrlsPaginatorOptions
	client    ListCrlsAPIClient
	params    *ListCrlsInput
	nextToken *string
	firstPage bool
}

// NewListCrlsPaginator returns a new ListCrlsPaginator
func NewListCrlsPaginator(client ListCrlsAPIClient, params *ListCrlsInput, optFns ...func(*ListCrlsPaginatorOptions)) *ListCrlsPaginator {
	if params == nil {
		params = &ListCrlsInput{}
	}

	options := ListCrlsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCrlsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCrlsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCrls page.
func (p *ListCrlsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCrlsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListCrls(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListCrls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rolesanywhere",
		OperationName: "ListCrls",
	}
}
