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

// ModifyRequest ModifyRequest
//
// swagger:model ModifyRequest
type ModifyRequest struct {

	// group ids
	GroupIds []strfmt.UUID `json:"groupIds"`
}

// Validate validates this modify request
func (m *ModifyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModifyRequest) validateGroupIds(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ModifyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModifyRequest) UnmarshalBinary(b []byte) error {
	var res ModifyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
