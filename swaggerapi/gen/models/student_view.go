// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StudentView student view
//
// student view
//
// swagger:model student_view
type StudentView struct {

	// id
	ID float64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this student view
func (m *StudentView) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this student view based on context it is used
func (m *StudentView) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StudentView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StudentView) UnmarshalBinary(b []byte) error {
	var res StudentView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
