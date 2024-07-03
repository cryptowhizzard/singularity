// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// RemoveWalletReader is a Reader for the RemoveWallet structure.
type RemoveWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveWalletNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveWalletInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /wallet/{address}] RemoveWallet", response, response.Code())
	}
}

// NewRemoveWalletNoContent creates a RemoveWalletNoContent with default headers values
func NewRemoveWalletNoContent() *RemoveWalletNoContent {
	return &RemoveWalletNoContent{}
}

/*
RemoveWalletNoContent describes a response with status code 204, with default header values.

No Content
*/
type RemoveWalletNoContent struct {
}

// IsSuccess returns true when this remove wallet no content response has a 2xx status code
func (o *RemoveWalletNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove wallet no content response has a 3xx status code
func (o *RemoveWalletNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove wallet no content response has a 4xx status code
func (o *RemoveWalletNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove wallet no content response has a 5xx status code
func (o *RemoveWalletNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this remove wallet no content response a status code equal to that given
func (o *RemoveWalletNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the remove wallet no content response
func (o *RemoveWalletNoContent) Code() int {
	return 204
}

func (o *RemoveWalletNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] removeWalletNoContent", 204)
}

func (o *RemoveWalletNoContent) String() string {
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] removeWalletNoContent", 204)
}

func (o *RemoveWalletNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveWalletBadRequest creates a RemoveWalletBadRequest with default headers values
func NewRemoveWalletBadRequest() *RemoveWalletBadRequest {
	return &RemoveWalletBadRequest{}
}

/*
RemoveWalletBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RemoveWalletBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this remove wallet bad request response has a 2xx status code
func (o *RemoveWalletBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove wallet bad request response has a 3xx status code
func (o *RemoveWalletBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove wallet bad request response has a 4xx status code
func (o *RemoveWalletBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove wallet bad request response has a 5xx status code
func (o *RemoveWalletBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove wallet bad request response a status code equal to that given
func (o *RemoveWalletBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove wallet bad request response
func (o *RemoveWalletBadRequest) Code() int {
	return 400
}

func (o *RemoveWalletBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] removeWalletBadRequest %s", 400, payload)
}

func (o *RemoveWalletBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] removeWalletBadRequest %s", 400, payload)
}

func (o *RemoveWalletBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RemoveWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveWalletInternalServerError creates a RemoveWalletInternalServerError with default headers values
func NewRemoveWalletInternalServerError() *RemoveWalletInternalServerError {
	return &RemoveWalletInternalServerError{}
}

/*
RemoveWalletInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RemoveWalletInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this remove wallet internal server error response has a 2xx status code
func (o *RemoveWalletInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove wallet internal server error response has a 3xx status code
func (o *RemoveWalletInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove wallet internal server error response has a 4xx status code
func (o *RemoveWalletInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove wallet internal server error response has a 5xx status code
func (o *RemoveWalletInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove wallet internal server error response a status code equal to that given
func (o *RemoveWalletInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove wallet internal server error response
func (o *RemoveWalletInternalServerError) Code() int {
	return 500
}

func (o *RemoveWalletInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] removeWalletInternalServerError %s", 500, payload)
}

func (o *RemoveWalletInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /wallet/{address}][%d] removeWalletInternalServerError %s", 500, payload)
}

func (o *RemoveWalletInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RemoveWalletInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
