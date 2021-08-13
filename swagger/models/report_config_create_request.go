// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportConfigCreateRequest ReportConfigCreateRequest
//
// swagger:model ReportConfigCreateRequest
type ReportConfigCreateRequest struct {

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// connector ids
	ConnectorIds []strfmt.UUID `json:"connectorIds"`

	// description
	Description string `json:"description,omitempty"`

	// format
	Format string `json:"format,omitempty"`

	// group ids
	GroupIds []strfmt.UUID `json:"groupIds"`

	// mandate Id
	MandateID string `json:"mandateId,omitempty"`

	// policies
	Policies []*PolicyFilter `json:"policies"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this report config create request
func (m *ReportConfigCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectorIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportConfigCreateRequest) validateConnectorIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectorIds) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectorIds); i++ {

		if err := validate.FormatOf("connectorIds"+"."+strconv.Itoa(i), "body", "uuid", m.ConnectorIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ReportConfigCreateRequest) validateGroupIds(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupIds) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupIds); i++ {

		if err := validate.FormatOf("groupIds"+"."+strconv.Itoa(i), "body", "uuid", m.GroupIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ReportConfigCreateRequest) validatePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportConfigCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportConfigCreateRequest) UnmarshalBinary(b []byte) error {
	var res ReportConfigCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
