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

// CreateAzureblobStorageReader is a Reader for the CreateAzureblobStorage structure.
type CreateAzureblobStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAzureblobStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAzureblobStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAzureblobStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAzureblobStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/azureblob] CreateAzureblobStorage", response, response.Code())
	}
}

// NewCreateAzureblobStorageOK creates a CreateAzureblobStorageOK with default headers values
func NewCreateAzureblobStorageOK() *CreateAzureblobStorageOK {
	return &CreateAzureblobStorageOK{}
}

/*
CreateAzureblobStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateAzureblobStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create azureblob storage o k response has a 2xx status code
func (o *CreateAzureblobStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create azureblob storage o k response has a 3xx status code
func (o *CreateAzureblobStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azureblob storage o k response has a 4xx status code
func (o *CreateAzureblobStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create azureblob storage o k response has a 5xx status code
func (o *CreateAzureblobStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create azureblob storage o k response a status code equal to that given
func (o *CreateAzureblobStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create azureblob storage o k response
func (o *CreateAzureblobStorageOK) Code() int {
	return 200
}

func (o *CreateAzureblobStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/azureblob][%d] createAzureblobStorageOK %s", 200, payload)
}

func (o *CreateAzureblobStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/azureblob][%d] createAzureblobStorageOK %s", 200, payload)
}

func (o *CreateAzureblobStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateAzureblobStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureblobStorageBadRequest creates a CreateAzureblobStorageBadRequest with default headers values
func NewCreateAzureblobStorageBadRequest() *CreateAzureblobStorageBadRequest {
	return &CreateAzureblobStorageBadRequest{}
}

/*
CreateAzureblobStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAzureblobStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create azureblob storage bad request response has a 2xx status code
func (o *CreateAzureblobStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create azureblob storage bad request response has a 3xx status code
func (o *CreateAzureblobStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azureblob storage bad request response has a 4xx status code
func (o *CreateAzureblobStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create azureblob storage bad request response has a 5xx status code
func (o *CreateAzureblobStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create azureblob storage bad request response a status code equal to that given
func (o *CreateAzureblobStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create azureblob storage bad request response
func (o *CreateAzureblobStorageBadRequest) Code() int {
	return 400
}

func (o *CreateAzureblobStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/azureblob][%d] createAzureblobStorageBadRequest %s", 400, payload)
}

func (o *CreateAzureblobStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/azureblob][%d] createAzureblobStorageBadRequest %s", 400, payload)
}

func (o *CreateAzureblobStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateAzureblobStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureblobStorageInternalServerError creates a CreateAzureblobStorageInternalServerError with default headers values
func NewCreateAzureblobStorageInternalServerError() *CreateAzureblobStorageInternalServerError {
	return &CreateAzureblobStorageInternalServerError{}
}

/*
CreateAzureblobStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAzureblobStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create azureblob storage internal server error response has a 2xx status code
func (o *CreateAzureblobStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create azureblob storage internal server error response has a 3xx status code
func (o *CreateAzureblobStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azureblob storage internal server error response has a 4xx status code
func (o *CreateAzureblobStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create azureblob storage internal server error response has a 5xx status code
func (o *CreateAzureblobStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create azureblob storage internal server error response a status code equal to that given
func (o *CreateAzureblobStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create azureblob storage internal server error response
func (o *CreateAzureblobStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateAzureblobStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/azureblob][%d] createAzureblobStorageInternalServerError %s", 500, payload)
}

func (o *CreateAzureblobStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/azureblob][%d] createAzureblobStorageInternalServerError %s", 500, payload)
}

func (o *CreateAzureblobStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateAzureblobStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
