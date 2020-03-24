// Code generated by go-swagger; DO NOT EDIT.

package wireless_settings

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

// NewUpdateNetworkWirelessSettingsParams creates a new UpdateNetworkWirelessSettingsParams object
// with the default values initialized.
func NewUpdateNetworkWirelessSettingsParams() *UpdateNetworkWirelessSettingsParams {
	var ()
	return &UpdateNetworkWirelessSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWirelessSettingsParamsWithTimeout creates a new UpdateNetworkWirelessSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkWirelessSettingsParamsWithTimeout(timeout time.Duration) *UpdateNetworkWirelessSettingsParams {
	var ()
	return &UpdateNetworkWirelessSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkWirelessSettingsParamsWithContext creates a new UpdateNetworkWirelessSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkWirelessSettingsParamsWithContext(ctx context.Context) *UpdateNetworkWirelessSettingsParams {
	var ()
	return &UpdateNetworkWirelessSettingsParams{

		Context: ctx,
	}
}

// NewUpdateNetworkWirelessSettingsParamsWithHTTPClient creates a new UpdateNetworkWirelessSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkWirelessSettingsParamsWithHTTPClient(client *http.Client) *UpdateNetworkWirelessSettingsParams {
	var ()
	return &UpdateNetworkWirelessSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkWirelessSettingsParams contains all the parameters to send to the API endpoint
for the update network wireless settings operation typically these are written to a http.Request
*/
type UpdateNetworkWirelessSettingsParams struct {

	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkWirelessSettings*/
	UpdateNetworkWirelessSettings *models.UpdateNetworkWirelessSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) WithTimeout(timeout time.Duration) *UpdateNetworkWirelessSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) WithContext(ctx context.Context) *UpdateNetworkWirelessSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) WithHTTPClient(client *http.Client) *UpdateNetworkWirelessSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) WithNetworkID(networkID string) *UpdateNetworkWirelessSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkWirelessSettings adds the updateNetworkWirelessSettings to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) WithUpdateNetworkWirelessSettings(updateNetworkWirelessSettings *models.UpdateNetworkWirelessSettings) *UpdateNetworkWirelessSettingsParams {
	o.SetUpdateNetworkWirelessSettings(updateNetworkWirelessSettings)
	return o
}

// SetUpdateNetworkWirelessSettings adds the updateNetworkWirelessSettings to the update network wireless settings params
func (o *UpdateNetworkWirelessSettingsParams) SetUpdateNetworkWirelessSettings(updateNetworkWirelessSettings *models.UpdateNetworkWirelessSettings) {
	o.UpdateNetworkWirelessSettings = updateNetworkWirelessSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWirelessSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkWirelessSettings != nil {
		if err := r.SetBodyParam(o.UpdateNetworkWirelessSettings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}