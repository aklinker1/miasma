// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/aklinker1/miasma/package/models"
)

// NewUpdateAppTraefikConfigParams creates a new UpdateAppTraefikConfigParams object
// with the default values initialized.
func NewUpdateAppTraefikConfigParams() *UpdateAppTraefikConfigParams {
	var ()
	return &UpdateAppTraefikConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppTraefikConfigParamsWithTimeout creates a new UpdateAppTraefikConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAppTraefikConfigParamsWithTimeout(timeout time.Duration) *UpdateAppTraefikConfigParams {
	var ()
	return &UpdateAppTraefikConfigParams{

		timeout: timeout,
	}
}

// NewUpdateAppTraefikConfigParamsWithContext creates a new UpdateAppTraefikConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAppTraefikConfigParamsWithContext(ctx context.Context) *UpdateAppTraefikConfigParams {
	var ()
	return &UpdateAppTraefikConfigParams{

		Context: ctx,
	}
}

// NewUpdateAppTraefikConfigParamsWithHTTPClient creates a new UpdateAppTraefikConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAppTraefikConfigParamsWithHTTPClient(client *http.Client) *UpdateAppTraefikConfigParams {
	var ()
	return &UpdateAppTraefikConfigParams{
		HTTPClient: client,
	}
}

/*UpdateAppTraefikConfigParams contains all the parameters to send to the API endpoint
for the update app traefik config operation typically these are written to a http.Request
*/
type UpdateAppTraefikConfigParams struct {

	/*AppID*/
	AppID strfmt.UUID4
	/*NewTraefikConfig*/
	NewTraefikConfig *models.InputTraefikPluginConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) WithTimeout(timeout time.Duration) *UpdateAppTraefikConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) WithContext(ctx context.Context) *UpdateAppTraefikConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) WithHTTPClient(client *http.Client) *UpdateAppTraefikConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) WithAppID(appID strfmt.UUID4) *UpdateAppTraefikConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) SetAppID(appID strfmt.UUID4) {
	o.AppID = appID
}

// WithNewTraefikConfig adds the newTraefikConfig to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) WithNewTraefikConfig(newTraefikConfig *models.InputTraefikPluginConfig) *UpdateAppTraefikConfigParams {
	o.SetNewTraefikConfig(newTraefikConfig)
	return o
}

// SetNewTraefikConfig adds the newTraefikConfig to the update app traefik config params
func (o *UpdateAppTraefikConfigParams) SetNewTraefikConfig(newTraefikConfig *models.InputTraefikPluginConfig) {
	o.NewTraefikConfig = newTraefikConfig
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppTraefikConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID.String()); err != nil {
		return err
	}

	if o.NewTraefikConfig != nil {
		if err := r.SetBodyParam(o.NewTraefikConfig); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}