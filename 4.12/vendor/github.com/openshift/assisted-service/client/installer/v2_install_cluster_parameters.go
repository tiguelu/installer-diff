// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewV2InstallClusterParams creates a new V2InstallClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2InstallClusterParams() *V2InstallClusterParams {
	return &V2InstallClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2InstallClusterParamsWithTimeout creates a new V2InstallClusterParams object
// with the ability to set a timeout on a request.
func NewV2InstallClusterParamsWithTimeout(timeout time.Duration) *V2InstallClusterParams {
	return &V2InstallClusterParams{
		timeout: timeout,
	}
}

// NewV2InstallClusterParamsWithContext creates a new V2InstallClusterParams object
// with the ability to set a context for a request.
func NewV2InstallClusterParamsWithContext(ctx context.Context) *V2InstallClusterParams {
	return &V2InstallClusterParams{
		Context: ctx,
	}
}

// NewV2InstallClusterParamsWithHTTPClient creates a new V2InstallClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2InstallClusterParamsWithHTTPClient(client *http.Client) *V2InstallClusterParams {
	return &V2InstallClusterParams{
		HTTPClient: client,
	}
}

/* V2InstallClusterParams contains all the parameters to send to the API endpoint
   for the v2 install cluster operation.

   Typically these are written to a http.Request.
*/
type V2InstallClusterParams struct {

	/* ClusterID.

	   The cluster to be installed.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 install cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2InstallClusterParams) WithDefaults() *V2InstallClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 install cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2InstallClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 install cluster params
func (o *V2InstallClusterParams) WithTimeout(timeout time.Duration) *V2InstallClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 install cluster params
func (o *V2InstallClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 install cluster params
func (o *V2InstallClusterParams) WithContext(ctx context.Context) *V2InstallClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 install cluster params
func (o *V2InstallClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 install cluster params
func (o *V2InstallClusterParams) WithHTTPClient(client *http.Client) *V2InstallClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 install cluster params
func (o *V2InstallClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the v2 install cluster params
func (o *V2InstallClusterParams) WithClusterID(clusterID strfmt.UUID) *V2InstallClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 install cluster params
func (o *V2InstallClusterParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *V2InstallClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
