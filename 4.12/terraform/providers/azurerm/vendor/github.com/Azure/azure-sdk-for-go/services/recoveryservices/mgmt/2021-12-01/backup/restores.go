package backup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RestoresClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type RestoresClient struct {
	BaseClient
}

// NewRestoresClient creates an instance of the RestoresClient client.
func NewRestoresClient(subscriptionID string) RestoresClient {
	return NewRestoresClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestoresClientWithBaseURI creates an instance of the RestoresClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRestoresClientWithBaseURI(baseURI string, subscriptionID string) RestoresClient {
	return RestoresClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Trigger restores the specified backed up data. This is an asynchronous operation. To know the status of this API
// call, use
// GetProtectedItemOperationResult API.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// fabricName - fabric name associated with the backed up items.
// containerName - container name associated with the backed up items.
// protectedItemName - backed up item to be restored.
// recoveryPointID - recovery point ID which represents the backed up data to be restored.
// parameters - resource restore request
func (client RestoresClient) Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters RestoreRequestResource) (result RestoresTriggerFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestoresClient.Trigger")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TriggerPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RestoresClient", "Trigger", nil, "Failure preparing request")
		return
	}

	result, err = client.TriggerSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RestoresClient", "Trigger", result.Response(), "Failure sending request")
		return
	}

	return
}

// TriggerPreparer prepares the Trigger request.
func (client RestoresClient) TriggerPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters RestoreRequestResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"recoveryPointId":   autorest.Encode("path", recoveryPointID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/restore", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TriggerSender sends the Trigger request. The method will close the
// http.Response Body if it receives an error.
func (client RestoresClient) TriggerSender(req *http.Request) (future RestoresTriggerFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// TriggerResponder handles the response to the Trigger request. The method always
// closes the http.Response Body.
func (client RestoresClient) TriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
