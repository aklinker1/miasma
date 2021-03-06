// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aklinker1/miasma/package/models"
)

// UpdateRunConfigReader is a Reader for the UpdateRunConfig structure.
type UpdateRunConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRunConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRunConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRunConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRunConfigOK creates a UpdateRunConfigOK with default headers values
func NewUpdateRunConfigOK() *UpdateRunConfigOK {
	return &UpdateRunConfigOK{}
}

/*UpdateRunConfigOK handles this case with default header values.

OK
*/
type UpdateRunConfigOK struct {
	Payload *models.RunConfig
}

func (o *UpdateRunConfigOK) Error() string {
	return fmt.Sprintf("[PUT /api/apps/{appName}/config][%d] updateRunConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateRunConfigOK) GetPayload() *models.RunConfig {
	return o.Payload
}

func (o *UpdateRunConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunConfigBadRequest creates a UpdateRunConfigBadRequest with default headers values
func NewUpdateRunConfigBadRequest() *UpdateRunConfigBadRequest {
	return &UpdateRunConfigBadRequest{}
}

/*UpdateRunConfigBadRequest handles this case with default header values.

Bad Request
*/
type UpdateRunConfigBadRequest struct {
	Payload string
}

func (o *UpdateRunConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/apps/{appName}/config][%d] updateRunConfigBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRunConfigBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateRunConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunConfigNotFound creates a UpdateRunConfigNotFound with default headers values
func NewUpdateRunConfigNotFound() *UpdateRunConfigNotFound {
	return &UpdateRunConfigNotFound{}
}

/*UpdateRunConfigNotFound handles this case with default header values.

Not Found
*/
type UpdateRunConfigNotFound struct {
	Payload string
}

func (o *UpdateRunConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/apps/{appName}/config][%d] updateRunConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRunConfigNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateRunConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunConfigDefault creates a UpdateRunConfigDefault with default headers values
func NewUpdateRunConfigDefault(code int) *UpdateRunConfigDefault {
	return &UpdateRunConfigDefault{
		_statusCode: code,
	}
}

/*UpdateRunConfigDefault handles this case with default header values.

Unknown Error
*/
type UpdateRunConfigDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the update run config default response
func (o *UpdateRunConfigDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRunConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /api/apps/{appName}/config][%d] updateRunConfig default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRunConfigDefault) GetPayload() string {
	return o.Payload
}

func (o *UpdateRunConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
