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

// NewRemoveAppTraefikConfigParams creates a new RemoveAppTraefikConfigParams object
// with the default values initialized.
func NewRemoveAppTraefikConfigParams() *RemoveAppTraefikConfigParams {
	var ()
	return &RemoveAppTraefikConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveAppTraefikConfigParamsWithTimeout creates a new RemoveAppTraefikConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveAppTraefikConfigParamsWithTimeout(timeout time.Duration) *RemoveAppTraefikConfigParams {
	var ()
	return &RemoveAppTraefikConfigParams{

		timeout: timeout,
	}
}

// NewRemoveAppTraefikConfigParamsWithContext creates a new RemoveAppTraefikConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveAppTraefikConfigParamsWithContext(ctx context.Context) *RemoveAppTraefikConfigParams {
	var ()
	return &RemoveAppTraefikConfigParams{

		Context: ctx,
	}
}

// NewRemoveAppTraefikConfigParamsWithHTTPClient creates a new RemoveAppTraefikConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveAppTraefikConfigParamsWithHTTPClient(client *http.Client) *RemoveAppTraefikConfigParams {
	var ()
	return &RemoveAppTraefikConfigParams{
		HTTPClient: client,
	}
}

/*RemoveAppTraefikConfigParams contains all the parameters to send to the API endpoint
for the remove app traefik config operation typically these are written to a http.Request
*/
type RemoveAppTraefikConfigParams struct {

	/*AppID*/
	AppID strfmt.UUID4

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove app traefik config params
func (o *RemoveAppTraefikConfigParams) WithTimeout(timeout time.Duration) *RemoveAppTraefikConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove app traefik config params
func (o *RemoveAppTraefikConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove app traefik config params
func (o *RemoveAppTraefikConfigParams) WithContext(ctx context.Context) *RemoveAppTraefikConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove app traefik config params
func (o *RemoveAppTraefikConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove app traefik config params
func (o *RemoveAppTraefikConfigParams) WithHTTPClient(client *http.Client) *RemoveAppTraefikConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove app traefik config params
func (o *RemoveAppTraefikConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the remove app traefik config params
func (o *RemoveAppTraefikConfigParams) WithAppID(appID strfmt.UUID4) *RemoveAppTraefikConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the remove app traefik config params
func (o *RemoveAppTraefikConfigParams) SetAppID(appID strfmt.UUID4) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveAppTraefikConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
