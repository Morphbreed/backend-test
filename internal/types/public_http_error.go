// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicHTTPError public Http error
//
// swagger:model publicHttpError
type PublicHTTPError struct {

	// More detailed, human-readable, optional explanation of the error
	Detail string `json:"detail,omitempty"`

	// HTTP status code returned for the error
	// Required: true
	// Maximum: 599
	// Minimum: 100
	Code *int64 `json:"status"`

	// Short, human-readable description of the error
	// Required: true
	Title *string `json:"title"`

	// Type of error returned, should be used for client-side error handling
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this public Http error
func (m *PublicHTTPError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicHTTPError) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Code); err != nil {
		return err
	}

	if err := validate.MinimumInt("status", "body", int64(*m.Code), 100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("status", "body", int64(*m.Code), 599, false); err != nil {
		return err
	}

	return nil
}

func (m *PublicHTTPError) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *PublicHTTPError) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicHTTPError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicHTTPError) UnmarshalBinary(b []byte) error {
	var res PublicHTTPError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
