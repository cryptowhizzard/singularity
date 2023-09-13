// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// CreateS3ChinaMobileStorageReader is a Reader for the CreateS3ChinaMobileStorage structure.
type CreateS3ChinaMobileStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3ChinaMobileStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3ChinaMobileStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3ChinaMobileStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3ChinaMobileStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/chinamobile] CreateS3ChinaMobileStorage", response, response.Code())
	}
}

// NewCreateS3ChinaMobileStorageOK creates a CreateS3ChinaMobileStorageOK with default headers values
func NewCreateS3ChinaMobileStorageOK() *CreateS3ChinaMobileStorageOK {
	return &CreateS3ChinaMobileStorageOK{}
}

/*
CreateS3ChinaMobileStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3ChinaMobileStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 china mobile storage o k response has a 2xx status code
func (o *CreateS3ChinaMobileStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 china mobile storage o k response has a 3xx status code
func (o *CreateS3ChinaMobileStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 china mobile storage o k response has a 4xx status code
func (o *CreateS3ChinaMobileStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 china mobile storage o k response has a 5xx status code
func (o *CreateS3ChinaMobileStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 china mobile storage o k response a status code equal to that given
func (o *CreateS3ChinaMobileStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 china mobile storage o k response
func (o *CreateS3ChinaMobileStorageOK) Code() int {
	return 200
}

func (o *CreateS3ChinaMobileStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/chinamobile][%d] createS3ChinaMobileStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3ChinaMobileStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/chinamobile][%d] createS3ChinaMobileStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3ChinaMobileStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3ChinaMobileStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3ChinaMobileStorageBadRequest creates a CreateS3ChinaMobileStorageBadRequest with default headers values
func NewCreateS3ChinaMobileStorageBadRequest() *CreateS3ChinaMobileStorageBadRequest {
	return &CreateS3ChinaMobileStorageBadRequest{}
}

/*
CreateS3ChinaMobileStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3ChinaMobileStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 china mobile storage bad request response has a 2xx status code
func (o *CreateS3ChinaMobileStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 china mobile storage bad request response has a 3xx status code
func (o *CreateS3ChinaMobileStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 china mobile storage bad request response has a 4xx status code
func (o *CreateS3ChinaMobileStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 china mobile storage bad request response has a 5xx status code
func (o *CreateS3ChinaMobileStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 china mobile storage bad request response a status code equal to that given
func (o *CreateS3ChinaMobileStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 china mobile storage bad request response
func (o *CreateS3ChinaMobileStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3ChinaMobileStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/chinamobile][%d] createS3ChinaMobileStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3ChinaMobileStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/chinamobile][%d] createS3ChinaMobileStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3ChinaMobileStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3ChinaMobileStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3ChinaMobileStorageInternalServerError creates a CreateS3ChinaMobileStorageInternalServerError with default headers values
func NewCreateS3ChinaMobileStorageInternalServerError() *CreateS3ChinaMobileStorageInternalServerError {
	return &CreateS3ChinaMobileStorageInternalServerError{}
}

/*
CreateS3ChinaMobileStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3ChinaMobileStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 china mobile storage internal server error response has a 2xx status code
func (o *CreateS3ChinaMobileStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 china mobile storage internal server error response has a 3xx status code
func (o *CreateS3ChinaMobileStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 china mobile storage internal server error response has a 4xx status code
func (o *CreateS3ChinaMobileStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 china mobile storage internal server error response has a 5xx status code
func (o *CreateS3ChinaMobileStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 china mobile storage internal server error response a status code equal to that given
func (o *CreateS3ChinaMobileStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 china mobile storage internal server error response
func (o *CreateS3ChinaMobileStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3ChinaMobileStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/chinamobile][%d] createS3ChinaMobileStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3ChinaMobileStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/chinamobile][%d] createS3ChinaMobileStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3ChinaMobileStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3ChinaMobileStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}