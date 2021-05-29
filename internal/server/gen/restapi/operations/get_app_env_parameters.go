// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetAppEnvParams creates a new GetAppEnvParams object
// no default values defined in spec.
func NewGetAppEnvParams() GetAppEnvParams {

	return GetAppEnvParams{}
}

// GetAppEnvParams contains all the bound params for the get app env operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAppEnv
type GetAppEnvParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*App name from the `-a|--app` flag
	  Required: true
	  In: path
	*/
	AppName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAppEnvParams() beforehand.
func (o *GetAppEnvParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAppName, rhkAppName, _ := route.Params.GetOK("appName")
	if err := o.bindAppName(rAppName, rhkAppName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAppName binds and validates parameter AppName from path.
func (o *GetAppEnvParams) bindAppName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AppName = raw

	return nil
}