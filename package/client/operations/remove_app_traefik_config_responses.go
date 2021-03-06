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

// RemoveAppTraefikConfigReader is a Reader for the RemoveAppTraefikConfig structure.
type RemoveAppTraefikConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAppTraefikConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveAppTraefikConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveAppTraefikConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRemoveAppTraefikConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveAppTraefikConfigOK creates a RemoveAppTraefikConfigOK with default headers values
func NewRemoveAppTraefikConfigOK() *RemoveAppTraefikConfigOK {
	return &RemoveAppTraefikConfigOK{}
}

/*RemoveAppTraefikConfigOK handles this case with default header values.

Created
*/
type RemoveAppTraefikConfigOK struct {
	Payload *models.TraefikPluginConfig
}

func (o *RemoveAppTraefikConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /api/plugins/traefik/{appId}][%d] removeAppTraefikConfigOK  %+v", 200, o.Payload)
}

func (o *RemoveAppTraefikConfigOK) GetPayload() *models.TraefikPluginConfig {
	return o.Payload
}

func (o *RemoveAppTraefikConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TraefikPluginConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAppTraefikConfigBadRequest creates a RemoveAppTraefikConfigBadRequest with default headers values
func NewRemoveAppTraefikConfigBadRequest() *RemoveAppTraefikConfigBadRequest {
	return &RemoveAppTraefikConfigBadRequest{}
}

/*RemoveAppTraefikConfigBadRequest handles this case with default header values.

Traefik plugin is not installed
*/
type RemoveAppTraefikConfigBadRequest struct {
	Payload string
}

func (o *RemoveAppTraefikConfigBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/plugins/traefik/{appId}][%d] removeAppTraefikConfigBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveAppTraefikConfigBadRequest) GetPayload() string {
	return o.Payload
}

func (o *RemoveAppTraefikConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAppTraefikConfigDefault creates a RemoveAppTraefikConfigDefault with default headers values
func NewRemoveAppTraefikConfigDefault(code int) *RemoveAppTraefikConfigDefault {
	return &RemoveAppTraefikConfigDefault{
		_statusCode: code,
	}
}

/*RemoveAppTraefikConfigDefault handles this case with default header values.

Unknown Error
*/
type RemoveAppTraefikConfigDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the remove app traefik config default response
func (o *RemoveAppTraefikConfigDefault) Code() int {
	return o._statusCode
}

func (o *RemoveAppTraefikConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/plugins/traefik/{appId}][%d] removeAppTraefikConfig default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveAppTraefikConfigDefault) GetPayload() string {
	return o.Payload
}

func (o *RemoveAppTraefikConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
