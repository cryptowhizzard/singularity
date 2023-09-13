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

// CreateS3IONOSStorageReader is a Reader for the CreateS3IONOSStorage structure.
type CreateS3IONOSStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3IONOSStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3IONOSStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3IONOSStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3IONOSStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/ionos] CreateS3IONOSStorage", response, response.Code())
	}
}

// NewCreateS3IONOSStorageOK creates a CreateS3IONOSStorageOK with default headers values
func NewCreateS3IONOSStorageOK() *CreateS3IONOSStorageOK {
	return &CreateS3IONOSStorageOK{}
}

/*
CreateS3IONOSStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3IONOSStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 i o n o s storage o k response has a 2xx status code
func (o *CreateS3IONOSStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 i o n o s storage o k response has a 3xx status code
func (o *CreateS3IONOSStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 i o n o s storage o k response has a 4xx status code
func (o *CreateS3IONOSStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 i o n o s storage o k response has a 5xx status code
func (o *CreateS3IONOSStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 i o n o s storage o k response a status code equal to that given
func (o *CreateS3IONOSStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 i o n o s storage o k response
func (o *CreateS3IONOSStorageOK) Code() int {
	return 200
}

func (o *CreateS3IONOSStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/ionos][%d] createS3IONOSStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3IONOSStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/ionos][%d] createS3IONOSStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3IONOSStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3IONOSStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3IONOSStorageBadRequest creates a CreateS3IONOSStorageBadRequest with default headers values
func NewCreateS3IONOSStorageBadRequest() *CreateS3IONOSStorageBadRequest {
	return &CreateS3IONOSStorageBadRequest{}
}

/*
CreateS3IONOSStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3IONOSStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 i o n o s storage bad request response has a 2xx status code
func (o *CreateS3IONOSStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 i o n o s storage bad request response has a 3xx status code
func (o *CreateS3IONOSStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 i o n o s storage bad request response has a 4xx status code
func (o *CreateS3IONOSStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 i o n o s storage bad request response has a 5xx status code
func (o *CreateS3IONOSStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 i o n o s storage bad request response a status code equal to that given
func (o *CreateS3IONOSStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 i o n o s storage bad request response
func (o *CreateS3IONOSStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3IONOSStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/ionos][%d] createS3IONOSStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3IONOSStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/ionos][%d] createS3IONOSStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3IONOSStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3IONOSStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3IONOSStorageInternalServerError creates a CreateS3IONOSStorageInternalServerError with default headers values
func NewCreateS3IONOSStorageInternalServerError() *CreateS3IONOSStorageInternalServerError {
	return &CreateS3IONOSStorageInternalServerError{}
}

/*
CreateS3IONOSStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3IONOSStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 i o n o s storage internal server error response has a 2xx status code
func (o *CreateS3IONOSStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 i o n o s storage internal server error response has a 3xx status code
func (o *CreateS3IONOSStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 i o n o s storage internal server error response has a 4xx status code
func (o *CreateS3IONOSStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 i o n o s storage internal server error response has a 5xx status code
func (o *CreateS3IONOSStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 i o n o s storage internal server error response a status code equal to that given
func (o *CreateS3IONOSStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 i o n o s storage internal server error response
func (o *CreateS3IONOSStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3IONOSStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/ionos][%d] createS3IONOSStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3IONOSStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/ionos][%d] createS3IONOSStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3IONOSStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3IONOSStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}