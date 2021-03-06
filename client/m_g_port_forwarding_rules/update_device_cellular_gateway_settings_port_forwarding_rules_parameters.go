// Code generated by go-swagger; DO NOT EDIT.

package m_g_port_forwarding_rules

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

// NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesParams creates a new UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams object
// with the default values initialized.
func NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesParams() *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	var ()
	return &UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesParamsWithTimeout creates a new UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesParamsWithTimeout(timeout time.Duration) *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	var ()
	return &UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams{

		timeout: timeout,
	}
}

// NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesParamsWithContext creates a new UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesParamsWithContext(ctx context.Context) *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	var ()
	return &UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams{

		Context: ctx,
	}
}

// NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesParamsWithHTTPClient creates a new UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeviceCellularGatewaySettingsPortForwardingRulesParamsWithHTTPClient(client *http.Client) *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	var ()
	return &UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams{
		HTTPClient: client,
	}
}

/*UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams contains all the parameters to send to the API endpoint
for the update device cellular gateway settings port forwarding rules operation typically these are written to a http.Request
*/
type UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams struct {

	/*Serial*/
	Serial string
	/*UpdateDeviceCellularGatewaySettingsPortForwardingRules*/
	UpdateDeviceCellularGatewaySettingsPortForwardingRules *models.UpdateDeviceCellularGatewaySettingsPortForwardingRules

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) WithTimeout(timeout time.Duration) *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) WithContext(ctx context.Context) *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) WithHTTPClient(client *http.Client) *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) WithSerial(serial string) *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceCellularGatewaySettingsPortForwardingRules adds the updateDeviceCellularGatewaySettingsPortForwardingRules to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) WithUpdateDeviceCellularGatewaySettingsPortForwardingRules(updateDeviceCellularGatewaySettingsPortForwardingRules *models.UpdateDeviceCellularGatewaySettingsPortForwardingRules) *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams {
	o.SetUpdateDeviceCellularGatewaySettingsPortForwardingRules(updateDeviceCellularGatewaySettingsPortForwardingRules)
	return o
}

// SetUpdateDeviceCellularGatewaySettingsPortForwardingRules adds the updateDeviceCellularGatewaySettingsPortForwardingRules to the update device cellular gateway settings port forwarding rules params
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) SetUpdateDeviceCellularGatewaySettingsPortForwardingRules(updateDeviceCellularGatewaySettingsPortForwardingRules *models.UpdateDeviceCellularGatewaySettingsPortForwardingRules) {
	o.UpdateDeviceCellularGatewaySettingsPortForwardingRules = updateDeviceCellularGatewaySettingsPortForwardingRules
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceCellularGatewaySettingsPortForwardingRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if o.UpdateDeviceCellularGatewaySettingsPortForwardingRules != nil {
		if err := r.SetBodyParam(o.UpdateDeviceCellularGatewaySettingsPortForwardingRules); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
