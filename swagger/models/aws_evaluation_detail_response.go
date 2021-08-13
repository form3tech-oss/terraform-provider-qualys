// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsEvaluationDetailResponse AwsEvaluationDetailResponse
//
// swagger:model AwsEvaluationDetailResponse
type AwsEvaluationDetailResponse struct {

	// account Id
	AccountID string `json:"accountId,omitempty"`

	// connector Id
	ConnectorID string `json:"connectorId,omitempty"`

	// evaluated on
	EvaluatedOn string `json:"evaluatedOn,omitempty"`

	// evaluation dates
	EvaluationDates *EvaluationDates `json:"evaluationDates,omitempty"`

	// evidences
	Evidences []*EvaluationEvidence `json:"evidences"`

	// region
	Region string `json:"region,omitempty"`

	// resource Id
	ResourceID string `json:"resourceId,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// result
	Result string `json:"result,omitempty"`
}

// Validate validates this aws evaluation detail response
func (m *AwsEvaluationDetailResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvaluationDates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvidences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsEvaluationDetailResponse) validateEvaluationDates(formats strfmt.Registry) error {

	if swag.IsZero(m.EvaluationDates) { // not required
		return nil
	}

	if m.EvaluationDates != nil {
		if err := m.EvaluationDates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluationDates")
			}
			return err
		}
	}

	return nil
}

func (m *AwsEvaluationDetailResponse) validateEvidences(formats strfmt.Registry) error {

	if swag.IsZero(m.Evidences) { // not required
		return nil
	}

	for i := 0; i < len(m.Evidences); i++ {
		if swag.IsZero(m.Evidences[i]) { // not required
			continue
		}

		if m.Evidences[i] != nil {
			if err := m.Evidences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evidences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsEvaluationDetailResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsEvaluationDetailResponse) UnmarshalBinary(b []byte) error {
	var res AwsEvaluationDetailResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
