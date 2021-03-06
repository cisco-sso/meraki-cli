// Code generated by go-swagger; DO NOT EDIT.

package s_m

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

	"github.com/cisco-sso/meraki-cli/models"
)

// NewUpdateNetworkSmAppPolarisParams creates a new UpdateNetworkSmAppPolarisParams object
// with the default values initialized.
func NewUpdateNetworkSmAppPolarisParams() *UpdateNetworkSmAppPolarisParams {
	var ()
	return &UpdateNetworkSmAppPolarisParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSmAppPolarisParamsWithTimeout creates a new UpdateNetworkSmAppPolarisParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkSmAppPolarisParamsWithTimeout(timeout time.Duration) *UpdateNetworkSmAppPolarisParams {
	var ()
	return &UpdateNetworkSmAppPolarisParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkSmAppPolarisParamsWithContext creates a new UpdateNetworkSmAppPolarisParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkSmAppPolarisParamsWithContext(ctx context.Context) *UpdateNetworkSmAppPolarisParams {
	var ()
	return &UpdateNetworkSmAppPolarisParams{

		Context: ctx,
	}
}

// NewUpdateNetworkSmAppPolarisParamsWithHTTPClient creates a new UpdateNetworkSmAppPolarisParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkSmAppPolarisParamsWithHTTPClient(client *http.Client) *UpdateNetworkSmAppPolarisParams {
	var ()
	return &UpdateNetworkSmAppPolarisParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkSmAppPolarisParams contains all the parameters to send to the API endpoint
for the update network sm app polaris operation typically these are written to a http.Request
*/
type UpdateNetworkSmAppPolarisParams struct {

	/*AppID*/
	AppID string
	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkSmAppPolaris*/
	UpdateNetworkSmAppPolaris *models.UpdateNetworkSmAppPolaris

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) WithTimeout(timeout time.Duration) *UpdateNetworkSmAppPolarisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) WithContext(ctx context.Context) *UpdateNetworkSmAppPolarisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) WithHTTPClient(client *http.Client) *UpdateNetworkSmAppPolarisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) WithAppID(appID string) *UpdateNetworkSmAppPolarisParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithNetworkID adds the networkID to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) WithNetworkID(networkID string) *UpdateNetworkSmAppPolarisParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkSmAppPolaris adds the updateNetworkSmAppPolaris to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) WithUpdateNetworkSmAppPolaris(updateNetworkSmAppPolaris *models.UpdateNetworkSmAppPolaris) *UpdateNetworkSmAppPolarisParams {
	o.SetUpdateNetworkSmAppPolaris(updateNetworkSmAppPolaris)
	return o
}

// SetUpdateNetworkSmAppPolaris adds the updateNetworkSmAppPolaris to the update network sm app polaris params
func (o *UpdateNetworkSmAppPolarisParams) SetUpdateNetworkSmAppPolaris(updateNetworkSmAppPolaris *models.UpdateNetworkSmAppPolaris) {
	o.UpdateNetworkSmAppPolaris = updateNetworkSmAppPolaris
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSmAppPolarisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkSmAppPolaris != nil {
		if err := r.SetBodyParam(o.UpdateNetworkSmAppPolaris); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
