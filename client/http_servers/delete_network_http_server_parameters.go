// Code generated by go-swagger; DO NOT EDIT.

package http_servers

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

// NewDeleteNetworkHTTPServerParams creates a new DeleteNetworkHTTPServerParams object
// with the default values initialized.
func NewDeleteNetworkHTTPServerParams() *DeleteNetworkHTTPServerParams {
	var ()
	return &DeleteNetworkHTTPServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkHTTPServerParamsWithTimeout creates a new DeleteNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNetworkHTTPServerParamsWithTimeout(timeout time.Duration) *DeleteNetworkHTTPServerParams {
	var ()
	return &DeleteNetworkHTTPServerParams{

		timeout: timeout,
	}
}

// NewDeleteNetworkHTTPServerParamsWithContext creates a new DeleteNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNetworkHTTPServerParamsWithContext(ctx context.Context) *DeleteNetworkHTTPServerParams {
	var ()
	return &DeleteNetworkHTTPServerParams{

		Context: ctx,
	}
}

// NewDeleteNetworkHTTPServerParamsWithHTTPClient creates a new DeleteNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNetworkHTTPServerParamsWithHTTPClient(client *http.Client) *DeleteNetworkHTTPServerParams {
	var ()
	return &DeleteNetworkHTTPServerParams{
		HTTPClient: client,
	}
}

/*DeleteNetworkHTTPServerParams contains all the parameters to send to the API endpoint
for the delete network Http server operation typically these are written to a http.Request
*/
type DeleteNetworkHTTPServerParams struct {

	/*ID*/
	ID string
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) WithTimeout(timeout time.Duration) *DeleteNetworkHTTPServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) WithContext(ctx context.Context) *DeleteNetworkHTTPServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) WithHTTPClient(client *http.Client) *DeleteNetworkHTTPServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) WithID(id string) *DeleteNetworkHTTPServerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) SetID(id string) {
	o.ID = id
}

// WithNetworkID adds the networkID to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) WithNetworkID(networkID string) *DeleteNetworkHTTPServerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network Http server params
func (o *DeleteNetworkHTTPServerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkHTTPServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
