// Code generated by go-swagger; DO NOT EDIT.

package file

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

// RetrieveFileReader is a Reader for the RetrieveFile structure.
type RetrieveFileReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveFileOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewRetrieveFilePartialContent(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRetrieveFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRetrieveFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /file/{id}/retrieve] RetrieveFile", response, response.Code())
	}
}

// NewRetrieveFileOK creates a RetrieveFileOK with default headers values
func NewRetrieveFileOK(writer io.Writer) *RetrieveFileOK {
	return &RetrieveFileOK{

		Payload: writer,
	}
}

/*
RetrieveFileOK describes a response with status code 200, with default header values.

OK
*/
type RetrieveFileOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this retrieve file o k response has a 2xx status code
func (o *RetrieveFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retrieve file o k response has a 3xx status code
func (o *RetrieveFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve file o k response has a 4xx status code
func (o *RetrieveFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve file o k response has a 5xx status code
func (o *RetrieveFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve file o k response a status code equal to that given
func (o *RetrieveFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the retrieve file o k response
func (o *RetrieveFileOK) Code() int {
	return 200
}

func (o *RetrieveFileOK) Error() string {
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFileOK", 200)
}

func (o *RetrieveFileOK) String() string {
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFileOK", 200)
}

func (o *RetrieveFileOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *RetrieveFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveFilePartialContent creates a RetrieveFilePartialContent with default headers values
func NewRetrieveFilePartialContent(writer io.Writer) *RetrieveFilePartialContent {
	return &RetrieveFilePartialContent{

		Payload: writer,
	}
}

/*
RetrieveFilePartialContent describes a response with status code 206, with default header values.

Partial Content
*/
type RetrieveFilePartialContent struct {
	Payload io.Writer
}

// IsSuccess returns true when this retrieve file partial content response has a 2xx status code
func (o *RetrieveFilePartialContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retrieve file partial content response has a 3xx status code
func (o *RetrieveFilePartialContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve file partial content response has a 4xx status code
func (o *RetrieveFilePartialContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve file partial content response has a 5xx status code
func (o *RetrieveFilePartialContent) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve file partial content response a status code equal to that given
func (o *RetrieveFilePartialContent) IsCode(code int) bool {
	return code == 206
}

// Code gets the status code for the retrieve file partial content response
func (o *RetrieveFilePartialContent) Code() int {
	return 206
}

func (o *RetrieveFilePartialContent) Error() string {
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFilePartialContent", 206)
}

func (o *RetrieveFilePartialContent) String() string {
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFilePartialContent", 206)
}

func (o *RetrieveFilePartialContent) GetPayload() io.Writer {
	return o.Payload
}

func (o *RetrieveFilePartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveFileBadRequest creates a RetrieveFileBadRequest with default headers values
func NewRetrieveFileBadRequest() *RetrieveFileBadRequest {
	return &RetrieveFileBadRequest{}
}

/*
RetrieveFileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RetrieveFileBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this retrieve file bad request response has a 2xx status code
func (o *RetrieveFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve file bad request response has a 3xx status code
func (o *RetrieveFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve file bad request response has a 4xx status code
func (o *RetrieveFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve file bad request response has a 5xx status code
func (o *RetrieveFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve file bad request response a status code equal to that given
func (o *RetrieveFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the retrieve file bad request response
func (o *RetrieveFileBadRequest) Code() int {
	return 400
}

func (o *RetrieveFileBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFileBadRequest %s", 400, payload)
}

func (o *RetrieveFileBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFileBadRequest %s", 400, payload)
}

func (o *RetrieveFileBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RetrieveFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveFileNotFound creates a RetrieveFileNotFound with default headers values
func NewRetrieveFileNotFound() *RetrieveFileNotFound {
	return &RetrieveFileNotFound{}
}

/*
RetrieveFileNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RetrieveFileNotFound struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this retrieve file not found response has a 2xx status code
func (o *RetrieveFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve file not found response has a 3xx status code
func (o *RetrieveFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve file not found response has a 4xx status code
func (o *RetrieveFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve file not found response has a 5xx status code
func (o *RetrieveFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve file not found response a status code equal to that given
func (o *RetrieveFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the retrieve file not found response
func (o *RetrieveFileNotFound) Code() int {
	return 404
}

func (o *RetrieveFileNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFileNotFound %s", 404, payload)
}

func (o *RetrieveFileNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFileNotFound %s", 404, payload)
}

func (o *RetrieveFileNotFound) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RetrieveFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveFileInternalServerError creates a RetrieveFileInternalServerError with default headers values
func NewRetrieveFileInternalServerError() *RetrieveFileInternalServerError {
	return &RetrieveFileInternalServerError{}
}

/*
RetrieveFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RetrieveFileInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this retrieve file internal server error response has a 2xx status code
func (o *RetrieveFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve file internal server error response has a 3xx status code
func (o *RetrieveFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve file internal server error response has a 4xx status code
func (o *RetrieveFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve file internal server error response has a 5xx status code
func (o *RetrieveFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this retrieve file internal server error response a status code equal to that given
func (o *RetrieveFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the retrieve file internal server error response
func (o *RetrieveFileInternalServerError) Code() int {
	return 500
}

func (o *RetrieveFileInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFileInternalServerError %s", 500, payload)
}

func (o *RetrieveFileInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}/retrieve][%d] retrieveFileInternalServerError %s", 500, payload)
}

func (o *RetrieveFileInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RetrieveFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
