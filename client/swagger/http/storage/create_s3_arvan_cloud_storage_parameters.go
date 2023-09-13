// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewCreateS3ArvanCloudStorageParams creates a new CreateS3ArvanCloudStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateS3ArvanCloudStorageParams() *CreateS3ArvanCloudStorageParams {
	return &CreateS3ArvanCloudStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateS3ArvanCloudStorageParamsWithTimeout creates a new CreateS3ArvanCloudStorageParams object
// with the ability to set a timeout on a request.
func NewCreateS3ArvanCloudStorageParamsWithTimeout(timeout time.Duration) *CreateS3ArvanCloudStorageParams {
	return &CreateS3ArvanCloudStorageParams{
		timeout: timeout,
	}
}

// NewCreateS3ArvanCloudStorageParamsWithContext creates a new CreateS3ArvanCloudStorageParams object
// with the ability to set a context for a request.
func NewCreateS3ArvanCloudStorageParamsWithContext(ctx context.Context) *CreateS3ArvanCloudStorageParams {
	return &CreateS3ArvanCloudStorageParams{
		Context: ctx,
	}
}

// NewCreateS3ArvanCloudStorageParamsWithHTTPClient creates a new CreateS3ArvanCloudStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateS3ArvanCloudStorageParamsWithHTTPClient(client *http.Client) *CreateS3ArvanCloudStorageParams {
	return &CreateS3ArvanCloudStorageParams{
		HTTPClient: client,
	}
}

/*
CreateS3ArvanCloudStorageParams contains all the parameters to send to the API endpoint

	for the create s3 arvan cloud storage operation.

	Typically these are written to a http.Request.
*/
type CreateS3ArvanCloudStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateS3ArvanCloudStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create s3 arvan cloud storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateS3ArvanCloudStorageParams) WithDefaults() *CreateS3ArvanCloudStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create s3 arvan cloud storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateS3ArvanCloudStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create s3 arvan cloud storage params
func (o *CreateS3ArvanCloudStorageParams) WithTimeout(timeout time.Duration) *CreateS3ArvanCloudStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create s3 arvan cloud storage params
func (o *CreateS3ArvanCloudStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create s3 arvan cloud storage params
func (o *CreateS3ArvanCloudStorageParams) WithContext(ctx context.Context) *CreateS3ArvanCloudStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create s3 arvan cloud storage params
func (o *CreateS3ArvanCloudStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create s3 arvan cloud storage params
func (o *CreateS3ArvanCloudStorageParams) WithHTTPClient(client *http.Client) *CreateS3ArvanCloudStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create s3 arvan cloud storage params
func (o *CreateS3ArvanCloudStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create s3 arvan cloud storage params
func (o *CreateS3ArvanCloudStorageParams) WithRequest(request *models.StorageCreateS3ArvanCloudStorageRequest) *CreateS3ArvanCloudStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create s3 arvan cloud storage params
func (o *CreateS3ArvanCloudStorageParams) SetRequest(request *models.StorageCreateS3ArvanCloudStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateS3ArvanCloudStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}