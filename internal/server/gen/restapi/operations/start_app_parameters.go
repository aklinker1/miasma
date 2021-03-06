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

// NewStartAppParams creates a new StartAppParams object
// no default values defined in spec.
func NewStartAppParams() StartAppParams {

	return StartAppParams{}
}

// StartAppParams contains all the bound params for the start app operation
// typically these are obtained from a http.Request
//
// swagger:parameters startApp
type StartAppParams struct {

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
// To ensure default values, the struct must have been initialized with NewStartAppParams() beforehand.
func (o *StartAppParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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
func (o *StartAppParams) bindAppName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AppName = raw

	return nil
}
