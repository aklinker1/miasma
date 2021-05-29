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

// NewEditAppParams creates a new EditAppParams object
// with the default values initialized.
func NewEditAppParams() *EditAppParams {
	var ()
	return &EditAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditAppParamsWithTimeout creates a new EditAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditAppParamsWithTimeout(timeout time.Duration) *EditAppParams {
	var ()
	return &EditAppParams{

		timeout: timeout,
	}
}

// NewEditAppParamsWithContext creates a new EditAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditAppParamsWithContext(ctx context.Context) *EditAppParams {
	var ()
	return &EditAppParams{

		Context: ctx,
	}
}

// NewEditAppParamsWithHTTPClient creates a new EditAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditAppParamsWithHTTPClient(client *http.Client) *EditAppParams {
	var ()
	return &EditAppParams{
		HTTPClient: client,
	}
}

/*EditAppParams contains all the parameters to send to the API endpoint
for the edit app operation typically these are written to a http.Request
*/
type EditAppParams struct {

	/*AppName
	  App name from the `-a|--app` flag

	*/
	AppName string
	/*NewApp*/
	NewApp *models.AppEdit

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit app params
func (o *EditAppParams) WithTimeout(timeout time.Duration) *EditAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit app params
func (o *EditAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit app params
func (o *EditAppParams) WithContext(ctx context.Context) *EditAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit app params
func (o *EditAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit app params
func (o *EditAppParams) WithHTTPClient(client *http.Client) *EditAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit app params
func (o *EditAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the edit app params
func (o *EditAppParams) WithAppName(appName string) *EditAppParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the edit app params
func (o *EditAppParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithNewApp adds the newApp to the edit app params
func (o *EditAppParams) WithNewApp(newApp *models.AppEdit) *EditAppParams {
	o.SetNewApp(newApp)
	return o
}

// SetNewApp adds the newApp to the edit app params
func (o *EditAppParams) SetNewApp(newApp *models.AppEdit) {
	o.NewApp = newApp
}

// WriteToRequest writes these params to a swagger request
func (o *EditAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	if o.NewApp != nil {
		if err := r.SetBodyParam(o.NewApp); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
