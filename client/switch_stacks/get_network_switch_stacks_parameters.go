// Code generated by go-swagger; DO NOT EDIT.

package switch_stacks

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

// NewGetNetworkSwitchStacksParams creates a new GetNetworkSwitchStacksParams object
// with the default values initialized.
func NewGetNetworkSwitchStacksParams() *GetNetworkSwitchStacksParams {
	var ()
	return &GetNetworkSwitchStacksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchStacksParamsWithTimeout creates a new GetNetworkSwitchStacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSwitchStacksParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchStacksParams {
	var ()
	return &GetNetworkSwitchStacksParams{

		timeout: timeout,
	}
}

// NewGetNetworkSwitchStacksParamsWithContext creates a new GetNetworkSwitchStacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSwitchStacksParamsWithContext(ctx context.Context) *GetNetworkSwitchStacksParams {
	var ()
	return &GetNetworkSwitchStacksParams{

		Context: ctx,
	}
}

// NewGetNetworkSwitchStacksParamsWithHTTPClient creates a new GetNetworkSwitchStacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSwitchStacksParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchStacksParams {
	var ()
	return &GetNetworkSwitchStacksParams{
		HTTPClient: client,
	}
}

/*GetNetworkSwitchStacksParams contains all the parameters to send to the API endpoint
for the get network switch stacks operation typically these are written to a http.Request
*/
type GetNetworkSwitchStacksParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network switch stacks params
func (o *GetNetworkSwitchStacksParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchStacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch stacks params
func (o *GetNetworkSwitchStacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch stacks params
func (o *GetNetworkSwitchStacksParams) WithContext(ctx context.Context) *GetNetworkSwitchStacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch stacks params
func (o *GetNetworkSwitchStacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch stacks params
func (o *GetNetworkSwitchStacksParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchStacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch stacks params
func (o *GetNetworkSwitchStacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network switch stacks params
func (o *GetNetworkSwitchStacksParams) WithNetworkID(networkID string) *GetNetworkSwitchStacksParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch stacks params
func (o *GetNetworkSwitchStacksParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchStacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
