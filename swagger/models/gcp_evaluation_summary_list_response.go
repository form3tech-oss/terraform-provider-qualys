// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GcpEvaluationSummaryListResponse GcpEvaluationSummaryListResponse
//
// swagger:model GcpEvaluationSummaryListResponse
type GcpEvaluationSummaryListResponse struct {

	// control Id
	ControlID string `json:"controlId,omitempty"`

	// control name
	ControlName string `json:"controlName,omitempty"`

	// criticality
	Criticality string `json:"criticality,omitempty"`

	// failed resources
	FailedResources int64 `json:"failedResources,omitempty"`

	// pass with exception resources
	PassWithExceptionResources int64 `json:"passWithExceptionResources,omitempty"`

	// passed resources
	PassedResources int64 `json:"passedResources,omitempty"`

	// policy names
	PolicyNames []string `json:"policyNames"`

	// result
	Result string `json:"result,omitempty"`

	// service
	Service string `json:"service,omitempty"`
}

// Validate validates this gcp evaluation summary list response
func (m *GcpEvaluationSummaryListResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpEvaluationSummaryListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpEvaluationSummaryListResponse) UnmarshalBinary(b []byte) error {
	var res GcpEvaluationSummaryListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
