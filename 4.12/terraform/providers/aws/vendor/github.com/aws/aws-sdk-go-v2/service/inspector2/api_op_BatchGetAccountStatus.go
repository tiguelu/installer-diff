// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the Amazon Inspector status of multiple Amazon Web Services accounts
// within your environment.
func (c *Client) BatchGetAccountStatus(ctx context.Context, params *BatchGetAccountStatusInput, optFns ...func(*Options)) (*BatchGetAccountStatusOutput, error) {
	if params == nil {
		params = &BatchGetAccountStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetAccountStatus", params, optFns, c.addOperationBatchGetAccountStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetAccountStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetAccountStatusInput struct {

	// The 12-digit Amazon Web Services account IDs of the accounts to retrieve Amazon
	// Inspector status for.
	AccountIds []string

	noSmithyDocumentSerde
}

type BatchGetAccountStatusOutput struct {

	// An array of objects that provide details on the status of Amazon Inspector for
	// each of the requested accounts.
	//
	// This member is required.
	Accounts []types.AccountState

	// An array of objects detailing any accounts that failed to enable Amazon
	// Inspector and why.
	FailedAccounts []types.FailedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetAccountStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetAccountStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetAccountStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetAccountStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetAccountStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "BatchGetAccountStatus",
	}
}
