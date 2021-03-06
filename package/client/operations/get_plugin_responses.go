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

// GetPluginReader is a Reader for the GetPlugin structure.
type GetPluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPluginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPluginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPluginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPluginOK creates a GetPluginOK with default headers values
func NewGetPluginOK() *GetPluginOK {
	return &GetPluginOK{}
}

/*GetPluginOK handles this case with default header values.

Created
*/
type GetPluginOK struct {
	Payload *models.Plugin
}

func (o *GetPluginOK) Error() string {
	return fmt.Sprintf("[GET /api/plugins/{pluginName}][%d] getPluginOK  %+v", 200, o.Payload)
}

func (o *GetPluginOK) GetPayload() *models.Plugin {
	return o.Payload
}

func (o *GetPluginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plugin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginNotFound creates a GetPluginNotFound with default headers values
func NewGetPluginNotFound() *GetPluginNotFound {
	return &GetPluginNotFound{}
}

/*GetPluginNotFound handles this case with default header values.

Not Found
*/
type GetPluginNotFound struct {
	Payload string
}

func (o *GetPluginNotFound) Error() string {
	return fmt.Sprintf("[GET /api/plugins/{pluginName}][%d] getPluginNotFound  %+v", 404, o.Payload)
}

func (o *GetPluginNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetPluginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginDefault creates a GetPluginDefault with default headers values
func NewGetPluginDefault(code int) *GetPluginDefault {
	return &GetPluginDefault{
		_statusCode: code,
	}
}

/*GetPluginDefault handles this case with default header values.

Unknown Error
*/
type GetPluginDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the get plugin default response
func (o *GetPluginDefault) Code() int {
	return o._statusCode
}

func (o *GetPluginDefault) Error() string {
	return fmt.Sprintf("[GET /api/plugins/{pluginName}][%d] getPlugin default  %+v", o._statusCode, o.Payload)
}

func (o *GetPluginDefault) GetPayload() string {
	return o.Payload
}

func (o *GetPluginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
