// Code generated by go-swagger; DO NOT EDIT.

package m_x_warm_spare_settings

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

// NewSwapNetworkWarmspareParams creates a new SwapNetworkWarmspareParams object
// with the default values initialized.
func NewSwapNetworkWarmspareParams() *SwapNetworkWarmspareParams {
	var ()
	return &SwapNetworkWarmspareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSwapNetworkWarmspareParamsWithTimeout creates a new SwapNetworkWarmspareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSwapNetworkWarmspareParamsWithTimeout(timeout time.Duration) *SwapNetworkWarmspareParams {
	var ()
	return &SwapNetworkWarmspareParams{

		timeout: timeout,
	}
}

// NewSwapNetworkWarmspareParamsWithContext creates a new SwapNetworkWarmspareParams object
// with the default values initialized, and the ability to set a context for a request
func NewSwapNetworkWarmspareParamsWithContext(ctx context.Context) *SwapNetworkWarmspareParams {
	var ()
	return &SwapNetworkWarmspareParams{

		Context: ctx,
	}
}

// NewSwapNetworkWarmspareParamsWithHTTPClient creates a new SwapNetworkWarmspareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSwapNetworkWarmspareParamsWithHTTPClient(client *http.Client) *SwapNetworkWarmspareParams {
	var ()
	return &SwapNetworkWarmspareParams{
		HTTPClient: client,
	}
}

/*SwapNetworkWarmspareParams contains all the parameters to send to the API endpoint
for the swap network warmspare operation typically these are written to a http.Request
*/
type SwapNetworkWarmspareParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the swap network warmspare params
func (o *SwapNetworkWarmspareParams) WithTimeout(timeout time.Duration) *SwapNetworkWarmspareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the swap network warmspare params
func (o *SwapNetworkWarmspareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the swap network warmspare params
func (o *SwapNetworkWarmspareParams) WithContext(ctx context.Context) *SwapNetworkWarmspareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the swap network warmspare params
func (o *SwapNetworkWarmspareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the swap network warmspare params
func (o *SwapNetworkWarmspareParams) WithHTTPClient(client *http.Client) *SwapNetworkWarmspareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the swap network warmspare params
func (o *SwapNetworkWarmspareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the swap network warmspare params
func (o *SwapNetworkWarmspareParams) WithNetworkID(networkID string) *SwapNetworkWarmspareParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the swap network warmspare params
func (o *SwapNetworkWarmspareParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *SwapNetworkWarmspareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
