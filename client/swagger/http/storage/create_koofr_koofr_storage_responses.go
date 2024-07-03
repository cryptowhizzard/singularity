// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// CreateKoofrKoofrStorageReader is a Reader for the CreateKoofrKoofrStorage structure.
type CreateKoofrKoofrStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateKoofrKoofrStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateKoofrKoofrStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateKoofrKoofrStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateKoofrKoofrStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/koofr/koofr] CreateKoofrKoofrStorage", response, response.Code())
	}
}

// NewCreateKoofrKoofrStorageOK creates a CreateKoofrKoofrStorageOK with default headers values
func NewCreateKoofrKoofrStorageOK() *CreateKoofrKoofrStorageOK {
	return &CreateKoofrKoofrStorageOK{}
}

/*
CreateKoofrKoofrStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateKoofrKoofrStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create koofr koofr storage o k response has a 2xx status code
func (o *CreateKoofrKoofrStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create koofr koofr storage o k response has a 3xx status code
func (o *CreateKoofrKoofrStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create koofr koofr storage o k response has a 4xx status code
func (o *CreateKoofrKoofrStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create koofr koofr storage o k response has a 5xx status code
func (o *CreateKoofrKoofrStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create koofr koofr storage o k response a status code equal to that given
func (o *CreateKoofrKoofrStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create koofr koofr storage o k response
func (o *CreateKoofrKoofrStorageOK) Code() int {
	return 200
}

func (o *CreateKoofrKoofrStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] createKoofrKoofrStorageOK %s", 200, payload)
}

func (o *CreateKoofrKoofrStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] createKoofrKoofrStorageOK %s", 200, payload)
}

func (o *CreateKoofrKoofrStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateKoofrKoofrStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateKoofrKoofrStorageBadRequest creates a CreateKoofrKoofrStorageBadRequest with default headers values
func NewCreateKoofrKoofrStorageBadRequest() *CreateKoofrKoofrStorageBadRequest {
	return &CreateKoofrKoofrStorageBadRequest{}
}

/*
CreateKoofrKoofrStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateKoofrKoofrStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create koofr koofr storage bad request response has a 2xx status code
func (o *CreateKoofrKoofrStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create koofr koofr storage bad request response has a 3xx status code
func (o *CreateKoofrKoofrStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create koofr koofr storage bad request response has a 4xx status code
func (o *CreateKoofrKoofrStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create koofr koofr storage bad request response has a 5xx status code
func (o *CreateKoofrKoofrStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create koofr koofr storage bad request response a status code equal to that given
func (o *CreateKoofrKoofrStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create koofr koofr storage bad request response
func (o *CreateKoofrKoofrStorageBadRequest) Code() int {
	return 400
}

func (o *CreateKoofrKoofrStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] createKoofrKoofrStorageBadRequest %s", 400, payload)
}

func (o *CreateKoofrKoofrStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] createKoofrKoofrStorageBadRequest %s", 400, payload)
}

func (o *CreateKoofrKoofrStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateKoofrKoofrStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateKoofrKoofrStorageInternalServerError creates a CreateKoofrKoofrStorageInternalServerError with default headers values
func NewCreateKoofrKoofrStorageInternalServerError() *CreateKoofrKoofrStorageInternalServerError {
	return &CreateKoofrKoofrStorageInternalServerError{}
}

/*
CreateKoofrKoofrStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateKoofrKoofrStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create koofr koofr storage internal server error response has a 2xx status code
func (o *CreateKoofrKoofrStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create koofr koofr storage internal server error response has a 3xx status code
func (o *CreateKoofrKoofrStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create koofr koofr storage internal server error response has a 4xx status code
func (o *CreateKoofrKoofrStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create koofr koofr storage internal server error response has a 5xx status code
func (o *CreateKoofrKoofrStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create koofr koofr storage internal server error response a status code equal to that given
func (o *CreateKoofrKoofrStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create koofr koofr storage internal server error response
func (o *CreateKoofrKoofrStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateKoofrKoofrStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] createKoofrKoofrStorageInternalServerError %s", 500, payload)
}

func (o *CreateKoofrKoofrStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/koofr/koofr][%d] createKoofrKoofrStorageInternalServerError %s", 500, payload)
}

func (o *CreateKoofrKoofrStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateKoofrKoofrStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
