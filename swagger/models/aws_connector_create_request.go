// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsConnectorCreateRequest AwsConnectorCreateRequest
//
// swagger:model AwsConnectorCreateRequest
type AwsConnectorCreateRequest struct {

	// arn
	Arn string `json:"arn,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// is china region
	IsChinaRegion bool `json:"isChinaRegion,omitempty"`

	// is gov cloud
	IsGovCloud bool `json:"isGovCloud,omitempty"`

	// is portal connector
	IsPortalConnector bool `json:"isPortalConnector,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// polling frequency
	PollingFrequency *RunFrequency `json:"pollingFrequency,omitempty"`

	// remediation enabled
	RemediationEnabled bool `json:"remediationEnabled,omitempty"`
}

// Validate validates this aws connector create request
func (m *AwsConnectorCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePollingFrequency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsConnectorCreateRequest) validatePollingFrequency(formats strfmt.Registry) error {

	if swag.IsZero(m.PollingFrequency) { // not required
		return nil
	}

	if m.PollingFrequency != nil {
		if err := m.PollingFrequency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pollingFrequency")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsConnectorCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsConnectorCreateRequest) UnmarshalBinary(b []byte) error {
	var res AwsConnectorCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
