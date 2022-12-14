// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelsNotificationTypeResponse models notification type response
//
// swagger:model models.NotificationTypeResponse
type ModelsNotificationTypeResponse struct {

	// default delivery channel
	DefaultDeliveryChannel string `json:"defaultDeliveryChannel,omitempty"`

	// default delivery schedule type
	DefaultDeliveryScheduleType string `json:"defaultDeliveryScheduleType,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// minimum seconds between notifications
	MinimumSecondsBetweenNotifications int64 `json:"minimumSecondsBetweenNotifications,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// qualifying action
	QualifyingAction string `json:"qualifyingAction,omitempty"`
}

// Validate validates this models notification type response
func (m *ModelsNotificationTypeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this models notification type response based on context it is used
func (m *ModelsNotificationTypeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsNotificationTypeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsNotificationTypeResponse) UnmarshalBinary(b []byte) error {
	var res ModelsNotificationTypeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
