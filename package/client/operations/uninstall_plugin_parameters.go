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
)

// NewUninstallPluginParams creates a new UninstallPluginParams object
// with the default values initialized.
func NewUninstallPluginParams() *UninstallPluginParams {
	var ()
	return &UninstallPluginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUninstallPluginParamsWithTimeout creates a new UninstallPluginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUninstallPluginParamsWithTimeout(timeout time.Duration) *UninstallPluginParams {
	var ()
	return &UninstallPluginParams{

		timeout: timeout,
	}
}

// NewUninstallPluginParamsWithContext creates a new UninstallPluginParams object
// with the default values initialized, and the ability to set a context for a request
func NewUninstallPluginParamsWithContext(ctx context.Context) *UninstallPluginParams {
	var ()
	return &UninstallPluginParams{

		Context: ctx,
	}
}

// NewUninstallPluginParamsWithHTTPClient creates a new UninstallPluginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUninstallPluginParamsWithHTTPClient(client *http.Client) *UninstallPluginParams {
	var ()
	return &UninstallPluginParams{
		HTTPClient: client,
	}
}

/*UninstallPluginParams contains all the parameters to send to the API endpoint
for the uninstall plugin operation typically these are written to a http.Request
*/
type UninstallPluginParams struct {

	/*PluginName*/
	PluginName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the uninstall plugin params
func (o *UninstallPluginParams) WithTimeout(timeout time.Duration) *UninstallPluginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the uninstall plugin params
func (o *UninstallPluginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the uninstall plugin params
func (o *UninstallPluginParams) WithContext(ctx context.Context) *UninstallPluginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the uninstall plugin params
func (o *UninstallPluginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the uninstall plugin params
func (o *UninstallPluginParams) WithHTTPClient(client *http.Client) *UninstallPluginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the uninstall plugin params
func (o *UninstallPluginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPluginName adds the pluginName to the uninstall plugin params
func (o *UninstallPluginParams) WithPluginName(pluginName string) *UninstallPluginParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the uninstall plugin params
func (o *UninstallPluginParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WriteToRequest writes these params to a swagger request
func (o *UninstallPluginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
