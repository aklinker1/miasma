// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/aklinker1/miasma/package/models"
)

// NewUpdateRunConfigParams creates a new UpdateRunConfigParams object
// no default values defined in spec.
func NewUpdateRunConfigParams() UpdateRunConfigParams {

	return UpdateRunConfigParams{}
}

// UpdateRunConfigParams contains all the bound params for the update run config operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateRunConfig
type UpdateRunConfigParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*App name from the `-a|--app` flag
	  Required: true
	  In: path
	*/
	AppName string
	/*
	  In: body
	*/
	NewRunConfig *models.InputRunConfig
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateRunConfigParams() beforehand.
func (o *UpdateRunConfigParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAppName, rhkAppName, _ := route.Params.GetOK("appName")
	if err := o.bindAppName(rAppName, rhkAppName, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.InputRunConfig
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("newRunConfig", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.NewRunConfig = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAppName binds and validates parameter AppName from path.
func (o *UpdateRunConfigParams) bindAppName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AppName = raw

	return nil
}