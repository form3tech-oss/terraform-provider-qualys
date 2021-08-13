// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// File File
//
// swagger:model File
type File struct {

	// absolute
	Absolute bool `json:"absolute,omitempty"`

	// absolute file
	AbsoluteFile *File `json:"absoluteFile,omitempty"`

	// absolute path
	AbsolutePath string `json:"absolutePath,omitempty"`

	// canonical file
	CanonicalFile *File `json:"canonicalFile,omitempty"`

	// canonical path
	CanonicalPath string `json:"canonicalPath,omitempty"`

	// directory
	Directory bool `json:"directory,omitempty"`

	// file
	File bool `json:"file,omitempty"`

	// free space
	FreeSpace int64 `json:"freeSpace,omitempty"`

	// hidden
	Hidden bool `json:"hidden,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parent
	Parent string `json:"parent,omitempty"`

	// parent file
	ParentFile *File `json:"parentFile,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// total space
	TotalSpace int64 `json:"totalSpace,omitempty"`

	// usable space
	UsableSpace int64 `json:"usableSpace,omitempty"`
}

// Validate validates this file
func (m *File) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbsoluteFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonicalFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *File) validateAbsoluteFile(formats strfmt.Registry) error {

	if swag.IsZero(m.AbsoluteFile) { // not required
		return nil
	}

	if m.AbsoluteFile != nil {
		if err := m.AbsoluteFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("absoluteFile")
			}
			return err
		}
	}

	return nil
}

func (m *File) validateCanonicalFile(formats strfmt.Registry) error {

	if swag.IsZero(m.CanonicalFile) { // not required
		return nil
	}

	if m.CanonicalFile != nil {
		if err := m.CanonicalFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("canonicalFile")
			}
			return err
		}
	}

	return nil
}

func (m *File) validateParentFile(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentFile) { // not required
		return nil
	}

	if m.ParentFile != nil {
		if err := m.ParentFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentFile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *File) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *File) UnmarshalBinary(b []byte) error {
	var res File
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
