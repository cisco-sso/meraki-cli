// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworkAccessPoliciesParams creates a new GetNetworkAccessPoliciesParams object
// with the default values initialized.
func NewGetNetworkAccessPoliciesParams() *GetNetworkAccessPoliciesParams {
	var ()
	return &GetNetworkAccessPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkAccessPoliciesParamsWithTimeout creates a new GetNetworkAccessPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkAccessPoliciesParamsWithTimeout(timeout time.Duration) *GetNetworkAccessPoliciesParams {
	var ()
	return &GetNetworkAccessPoliciesParams{

		timeout: timeout,
	}
}

// NewGetNetworkAccessPoliciesParamsWithContext creates a new GetNetworkAccessPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkAccessPoliciesParamsWithContext(ctx context.Context) *GetNetworkAccessPoliciesParams {
	var ()
	return &GetNetworkAccessPoliciesParams{

		Context: ctx,
	}
}

// NewGetNetworkAccessPoliciesParamsWithHTTPClient creates a new GetNetworkAccessPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkAccessPoliciesParamsWithHTTPClient(client *http.Client) *GetNetworkAccessPoliciesParams {
	var ()
	return &GetNetworkAccessPoliciesParams{
		HTTPClient: client,
	}
}

/*GetNetworkAccessPoliciesParams contains all the parameters to send to the API endpoint
for the get network access policies operation typically these are written to a http.Request
*/
type GetNetworkAccessPoliciesParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network access policies params
func (o *GetNetworkAccessPoliciesParams) WithTimeout(timeout time.Duration) *GetNetworkAccessPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network access policies params
func (o *GetNetworkAccessPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network access policies params
func (o *GetNetworkAccessPoliciesParams) WithContext(ctx context.Context) *GetNetworkAccessPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network access policies params
func (o *GetNetworkAccessPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network access policies params
func (o *GetNetworkAccessPoliciesParams) WithHTTPClient(client *http.Client) *GetNetworkAccessPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network access policies params
func (o *GetNetworkAccessPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network access policies params
func (o *GetNetworkAccessPoliciesParams) WithNetworkID(networkID string) *GetNetworkAccessPoliciesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network access policies params
func (o *GetNetworkAccessPoliciesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkAccessPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}