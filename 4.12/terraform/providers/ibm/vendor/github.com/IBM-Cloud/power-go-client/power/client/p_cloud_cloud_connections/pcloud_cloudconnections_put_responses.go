// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsPutReader is a Reader for the PcloudCloudconnectionsPut structure.
type PcloudCloudconnectionsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPcloudCloudconnectionsPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsPutRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudCloudconnectionsPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudCloudconnectionsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudconnectionsPutOK creates a PcloudCloudconnectionsPutOK with default headers values
func NewPcloudCloudconnectionsPutOK() *PcloudCloudconnectionsPutOK {
	return &PcloudCloudconnectionsPutOK{}
}

/* PcloudCloudconnectionsPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsPutOK struct {
	Payload *models.CloudConnection
}

func (o *PcloudCloudconnectionsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudconnectionsPutOK) GetPayload() *models.CloudConnection {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutAccepted creates a PcloudCloudconnectionsPutAccepted with default headers values
func NewPcloudCloudconnectionsPutAccepted() *PcloudCloudconnectionsPutAccepted {
	return &PcloudCloudconnectionsPutAccepted{}
}

/* PcloudCloudconnectionsPutAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsPutAccepted struct {
	Payload *models.JobReference
}

func (o *PcloudCloudconnectionsPutAccepted) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutAccepted  %+v", 202, o.Payload)
}
func (o *PcloudCloudconnectionsPutAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutBadRequest creates a PcloudCloudconnectionsPutBadRequest with default headers values
func NewPcloudCloudconnectionsPutBadRequest() *PcloudCloudconnectionsPutBadRequest {
	return &PcloudCloudconnectionsPutBadRequest{}
}

/* PcloudCloudconnectionsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudconnectionsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutUnauthorized creates a PcloudCloudconnectionsPutUnauthorized with default headers values
func NewPcloudCloudconnectionsPutUnauthorized() *PcloudCloudconnectionsPutUnauthorized {
	return &PcloudCloudconnectionsPutUnauthorized{}
}

/* PcloudCloudconnectionsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsPutUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudconnectionsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutNotFound creates a PcloudCloudconnectionsPutNotFound with default headers values
func NewPcloudCloudconnectionsPutNotFound() *PcloudCloudconnectionsPutNotFound {
	return &PcloudCloudconnectionsPutNotFound{}
}

/* PcloudCloudconnectionsPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsPutNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutNotFound  %+v", 404, o.Payload)
}
func (o *PcloudCloudconnectionsPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutMethodNotAllowed creates a PcloudCloudconnectionsPutMethodNotAllowed with default headers values
func NewPcloudCloudconnectionsPutMethodNotAllowed() *PcloudCloudconnectionsPutMethodNotAllowed {
	return &PcloudCloudconnectionsPutMethodNotAllowed{}
}

/* PcloudCloudconnectionsPutMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type PcloudCloudconnectionsPutMethodNotAllowed struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *PcloudCloudconnectionsPutMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutRequestTimeout creates a PcloudCloudconnectionsPutRequestTimeout with default headers values
func NewPcloudCloudconnectionsPutRequestTimeout() *PcloudCloudconnectionsPutRequestTimeout {
	return &PcloudCloudconnectionsPutRequestTimeout{}
}

/* PcloudCloudconnectionsPutRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsPutRequestTimeout struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPutRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutRequestTimeout  %+v", 408, o.Payload)
}
func (o *PcloudCloudconnectionsPutRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutConflict creates a PcloudCloudconnectionsPutConflict with default headers values
func NewPcloudCloudconnectionsPutConflict() *PcloudCloudconnectionsPutConflict {
	return &PcloudCloudconnectionsPutConflict{}
}

/* PcloudCloudconnectionsPutConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudCloudconnectionsPutConflict struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPutConflict) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutConflict  %+v", 409, o.Payload)
}
func (o *PcloudCloudconnectionsPutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutUnprocessableEntity creates a PcloudCloudconnectionsPutUnprocessableEntity with default headers values
func NewPcloudCloudconnectionsPutUnprocessableEntity() *PcloudCloudconnectionsPutUnprocessableEntity {
	return &PcloudCloudconnectionsPutUnprocessableEntity{}
}

/* PcloudCloudconnectionsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudconnectionsPutUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudCloudconnectionsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutInternalServerError creates a PcloudCloudconnectionsPutInternalServerError with default headers values
func NewPcloudCloudconnectionsPutInternalServerError() *PcloudCloudconnectionsPutInternalServerError {
	return &PcloudCloudconnectionsPutInternalServerError{}
}

/* PcloudCloudconnectionsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudconnectionsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
