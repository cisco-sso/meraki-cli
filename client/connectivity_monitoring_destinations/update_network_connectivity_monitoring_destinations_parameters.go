// Code generated by go-swagger; DO NOT EDIT.

package connectivity_monitoring_destinations

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

// NewUpdateNetworkConnectivityMonitoringDestinationsParams creates a new UpdateNetworkConnectivityMonitoringDestinationsParams object
// with the default values initialized.
func NewUpdateNetworkConnectivityMonitoringDestinationsParams() *UpdateNetworkConnectivityMonitoringDestinationsParams {
	var ()
	return &UpdateNetworkConnectivityMonitoringDestinationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkConnectivityMonitoringDestinationsParamsWithTimeout creates a new UpdateNetworkConnectivityMonitoringDestinationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkConnectivityMonitoringDestinationsParamsWithTimeout(timeout time.Duration) *UpdateNetworkConnectivityMonitoringDestinationsParams {
	var ()
	return &UpdateNetworkConnectivityMonitoringDestinationsParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkConnectivityMonitoringDestinationsParamsWithContext creates a new UpdateNetworkConnectivityMonitoringDestinationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkConnectivityMonitoringDestinationsParamsWithContext(ctx context.Context) *UpdateNetworkConnectivityMonitoringDestinationsParams {
	var ()
	return &UpdateNetworkConnectivityMonitoringDestinationsParams{

		Context: ctx,
	}
}

// NewUpdateNetworkConnectivityMonitoringDestinationsParamsWithHTTPClient creates a new UpdateNetworkConnectivityMonitoringDestinationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkConnectivityMonitoringDestinationsParamsWithHTTPClient(client *http.Client) *UpdateNetworkConnectivityMonitoringDestinationsParams {
	var ()
	return &UpdateNetworkConnectivityMonitoringDestinationsParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkConnectivityMonitoringDestinationsParams contains all the parameters to send to the API endpoint
for the update network connectivity monitoring destinations operation typically these are written to a http.Request
*/
type UpdateNetworkConnectivityMonitoringDestinationsParams struct {

	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkConnectivityMonitoringDestinations*/
	UpdateNetworkConnectivityMonitoringDestinations *models.UpdateNetworkConnectivityMonitoringDestinations

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) WithTimeout(timeout time.Duration) *UpdateNetworkConnectivityMonitoringDestinationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) WithContext(ctx context.Context) *UpdateNetworkConnectivityMonitoringDestinationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) WithHTTPClient(client *http.Client) *UpdateNetworkConnectivityMonitoringDestinationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) WithNetworkID(networkID string) *UpdateNetworkConnectivityMonitoringDestinationsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkConnectivityMonitoringDestinations adds the updateNetworkConnectivityMonitoringDestinations to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) WithUpdateNetworkConnectivityMonitoringDestinations(updateNetworkConnectivityMonitoringDestinations *models.UpdateNetworkConnectivityMonitoringDestinations) *UpdateNetworkConnectivityMonitoringDestinationsParams {
	o.SetUpdateNetworkConnectivityMonitoringDestinations(updateNetworkConnectivityMonitoringDestinations)
	return o
}

// SetUpdateNetworkConnectivityMonitoringDestinations adds the updateNetworkConnectivityMonitoringDestinations to the update network connectivity monitoring destinations params
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) SetUpdateNetworkConnectivityMonitoringDestinations(updateNetworkConnectivityMonitoringDestinations *models.UpdateNetworkConnectivityMonitoringDestinations) {
	o.UpdateNetworkConnectivityMonitoringDestinations = updateNetworkConnectivityMonitoringDestinations
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkConnectivityMonitoringDestinationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkConnectivityMonitoringDestinations != nil {
		if err := r.SetBodyParam(o.UpdateNetworkConnectivityMonitoringDestinations); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
