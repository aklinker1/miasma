// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aklinker1/miasma/package/models"
)

// EditAppOKCode is the HTTP code returned for type EditAppOK
const EditAppOKCode int = 200

/*EditAppOK OK

swagger:response editAppOK
*/
type EditAppOK struct {

	/*
	  In: Body
	*/
	Payload *models.App `json:"body,omitempty"`
}

// NewEditAppOK creates EditAppOK with default headers values
func NewEditAppOK() *EditAppOK {

	return &EditAppOK{}
}

// WithPayload adds the payload to the edit app o k response
func (o *EditAppOK) WithPayload(payload *models.App) *EditAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit app o k response
func (o *EditAppOK) SetPayload(payload *models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EditAppBadRequestCode is the HTTP code returned for type EditAppBadRequest
const EditAppBadRequestCode int = 400

/*EditAppBadRequest BadRequest

swagger:response editAppBadRequest
*/
type EditAppBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewEditAppBadRequest creates EditAppBadRequest with default headers values
func NewEditAppBadRequest() *EditAppBadRequest {

	return &EditAppBadRequest{}
}

// WithPayload adds the payload to the edit app bad request response
func (o *EditAppBadRequest) WithPayload(payload string) *EditAppBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit app bad request response
func (o *EditAppBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditAppBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// EditAppNotFoundCode is the HTTP code returned for type EditAppNotFound
const EditAppNotFoundCode int = 404

/*EditAppNotFound Not Found

swagger:response editAppNotFound
*/
type EditAppNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewEditAppNotFound creates EditAppNotFound with default headers values
func NewEditAppNotFound() *EditAppNotFound {

	return &EditAppNotFound{}
}

// WithPayload adds the payload to the edit app not found response
func (o *EditAppNotFound) WithPayload(payload string) *EditAppNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit app not found response
func (o *EditAppNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditAppNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*EditAppDefault Unknown Error

swagger:response editAppDefault
*/
type EditAppDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewEditAppDefault creates EditAppDefault with default headers values
func NewEditAppDefault(code int) *EditAppDefault {
	if code <= 0 {
		code = 500
	}

	return &EditAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the edit app default response
func (o *EditAppDefault) WithStatusCode(code int) *EditAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the edit app default response
func (o *EditAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the edit app default response
func (o *EditAppDefault) WithPayload(payload string) *EditAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit app default response
func (o *EditAppDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
