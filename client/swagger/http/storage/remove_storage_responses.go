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

// RemoveStorageReader is a Reader for the RemoveStorage structure.
type RemoveStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveStorageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /storage/{name}] RemoveStorage", response, response.Code())
	}
}

// NewRemoveStorageNoContent creates a RemoveStorageNoContent with default headers values
func NewRemoveStorageNoContent() *RemoveStorageNoContent {
	return &RemoveStorageNoContent{}
}

/*
RemoveStorageNoContent describes a response with status code 204, with default header values.

No Content
*/
type RemoveStorageNoContent struct {
}

// IsSuccess returns true when this remove storage no content response has a 2xx status code
func (o *RemoveStorageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove storage no content response has a 3xx status code
func (o *RemoveStorageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove storage no content response has a 4xx status code
func (o *RemoveStorageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove storage no content response has a 5xx status code
func (o *RemoveStorageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this remove storage no content response a status code equal to that given
func (o *RemoveStorageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the remove storage no content response
func (o *RemoveStorageNoContent) Code() int {
	return 204
}

func (o *RemoveStorageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /storage/{name}][%d] removeStorageNoContent", 204)
}

func (o *RemoveStorageNoContent) String() string {
	return fmt.Sprintf("[DELETE /storage/{name}][%d] removeStorageNoContent", 204)
}

func (o *RemoveStorageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveStorageBadRequest creates a RemoveStorageBadRequest with default headers values
func NewRemoveStorageBadRequest() *RemoveStorageBadRequest {
	return &RemoveStorageBadRequest{}
}

/*
RemoveStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RemoveStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this remove storage bad request response has a 2xx status code
func (o *RemoveStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove storage bad request response has a 3xx status code
func (o *RemoveStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove storage bad request response has a 4xx status code
func (o *RemoveStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove storage bad request response has a 5xx status code
func (o *RemoveStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove storage bad request response a status code equal to that given
func (o *RemoveStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove storage bad request response
func (o *RemoveStorageBadRequest) Code() int {
	return 400
}

func (o *RemoveStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/{name}][%d] removeStorageBadRequest %s", 400, payload)
}

func (o *RemoveStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/{name}][%d] removeStorageBadRequest %s", 400, payload)
}

func (o *RemoveStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RemoveStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveStorageInternalServerError creates a RemoveStorageInternalServerError with default headers values
func NewRemoveStorageInternalServerError() *RemoveStorageInternalServerError {
	return &RemoveStorageInternalServerError{}
}

/*
RemoveStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RemoveStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this remove storage internal server error response has a 2xx status code
func (o *RemoveStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove storage internal server error response has a 3xx status code
func (o *RemoveStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove storage internal server error response has a 4xx status code
func (o *RemoveStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove storage internal server error response has a 5xx status code
func (o *RemoveStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove storage internal server error response a status code equal to that given
func (o *RemoveStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove storage internal server error response
func (o *RemoveStorageInternalServerError) Code() int {
	return 500
}

func (o *RemoveStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/{name}][%d] removeStorageInternalServerError %s", 500, payload)
}

func (o *RemoveStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/{name}][%d] removeStorageInternalServerError %s", 500, payload)
}

func (o *RemoveStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RemoveStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
