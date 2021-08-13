// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SlackUpdateRequest SlackUpdateRequest
//
// swagger:model SlackUpdateRequest
type SlackUpdateRequest struct {

	// action description
	ActionDescription string `json:"actionDescription,omitempty"`

	// action name
	ActionName string `json:"actionName,omitempty"`

	// channel
	Channel string `json:"channel,omitempty"`

	// webhook Uri
	WebhookURI string `json:"webhookUri,omitempty"`
}

// Validate validates this slack update request
func (m *SlackUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SlackUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlackUpdateRequest) UnmarshalBinary(b []byte) error {
	var res SlackUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
