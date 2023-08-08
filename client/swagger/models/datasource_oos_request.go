// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatasourceOosRequest datasource oos request
//
// swagger:model datasource.OosRequest
type DatasourceOosRequest struct {

	// Chunk size to use for uploading.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Object storage compartment OCID
	Compartment string `json:"compartment,omitempty"`

	// Path to OCI config file
	ConfigFile *string `json:"configFile,omitempty"`

	// Profile name inside the oci config file
	ConfigProfile *string `json:"configProfile,omitempty"`

	// Cutoff for switching to multipart copy.
	CopyCutoff *string `json:"copyCutoff,omitempty"`

	// Timeout for copy.
	CopyTimeout *string `json:"copyTimeout,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// Don't store MD5 checksum with object metadata.
	DisableChecksum *string `json:"disableChecksum,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint for Object storage API.
	Endpoint string `json:"endpoint,omitempty"`

	// If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	LeavePartsOnError *string `json:"leavePartsOnError,omitempty"`

	// Object storage namespace
	Namespace string `json:"namespace,omitempty"`

	// If set, don't attempt to check the bucket exists or create it.
	NoCheckBucket *string `json:"noCheckBucket,omitempty"`

	// Choose your Auth Provider
	Provider *string `json:"provider,omitempty"`

	// Object storage Region
	Region string `json:"region,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm,omitempty"`

	// To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKey string `json:"sseCustomerKey,omitempty"`

	// To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKeyFile string `json:"sseCustomerKeyFile,omitempty"`

	// If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256,omitempty"`

	// if using using your own master key in vault, this header specifies the
	SseKmsKeyID string `json:"sseKmsKeyId,omitempty"`

	// The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	StorageTier *string `json:"storageTier,omitempty"`

	// Concurrency for multipart uploads.
	UploadConcurrency *string `json:"uploadConcurrency,omitempty"`

	// Cutoff for switching to chunked upload.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`
}

// Validate validates this datasource oos request
func (m *DatasourceOosRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteAfterExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescanInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanningState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceOosRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceOosRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceOosRequest) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceOosRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource oos request based on the context it is used
func (m *DatasourceOosRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceOosRequest) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceOosRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceOosRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceOosRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}