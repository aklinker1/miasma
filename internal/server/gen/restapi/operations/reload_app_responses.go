// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ReloadAppCreatedCode is the HTTP code returned for type ReloadAppCreated
const ReloadAppCreatedCode int = 201

/*ReloadAppCreated Reloaded

swagger:response reloadAppCreated
*/
type ReloadAppCreated struct {
}

// NewReloadAppCreated creates ReloadAppCreated with default headers values
func NewReloadAppCreated() *ReloadAppCreated {

	return &ReloadAppCreated{}
}

// WriteResponse to the client
func (o *ReloadAppCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// ReloadAppNotFoundCode is the HTTP code returned for type ReloadAppNotFound
const ReloadAppNotFoundCode int = 404

/*ReloadAppNotFound Not Found

swagger:response reloadAppNotFound
*/
type ReloadAppNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewReloadAppNotFound creates ReloadAppNotFound with default headers values
func NewReloadAppNotFound() *ReloadAppNotFound {

	return &ReloadAppNotFound{}
}

// WithPayload adds the payload to the reload app not found response
func (o *ReloadAppNotFound) WithPayload(payload string) *ReloadAppNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reload app not found response
func (o *ReloadAppNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReloadAppNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*ReloadAppDefault Unknown Error

swagger:response reloadAppDefault
*/
type ReloadAppDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewReloadAppDefault creates ReloadAppDefault with default headers values
func NewReloadAppDefault(code int) *ReloadAppDefault {
	if code <= 0 {
		code = 500
	}

	return &ReloadAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the reload app default response
func (o *ReloadAppDefault) WithStatusCode(code int) *ReloadAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the reload app default response
func (o *ReloadAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the reload app default response
func (o *ReloadAppDefault) WithPayload(payload string) *ReloadAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reload app default response
func (o *ReloadAppDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReloadAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}