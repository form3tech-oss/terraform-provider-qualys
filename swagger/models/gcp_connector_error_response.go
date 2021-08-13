// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GcpConnectorErrorResponse GcpConnectorErrorResponse
//
// swagger:model GcpConnectorErrorResponse
type GcpConnectorErrorResponse struct {

	// connector Id
	ConnectorID string `json:"connectorId,omitempty"`

	// connector name
	ConnectorName string `json:"connectorName,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// occured on
	OccuredOn string `json:"occuredOn,omitempty"`

	// region
	Region string `json:"region,omitempty"`
}

// Validate validates this gcp connector error response
func (m *GcpConnectorErrorResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpConnectorErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpConnectorErrorResponse) UnmarshalBinary(b []byte) error {
	var res GcpConnectorErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
