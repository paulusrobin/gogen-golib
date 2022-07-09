package validator

import (
	"errors"
)

// ValidationError error object returned from validation.
type ValidationError struct {
	Message string   `json:"message" xml:"message"`
	Details []string `json:"details,omitempty" xml:"details,omitempty"`
	detail  map[string]string
}

// Error function implement error interface.
func (ve ValidationError) Error() string {
	return ve.Message
}

// Detail function to get map detail of ValidationError.
func (ve ValidationError) Detail() map[string]string {
	return ve.detail
}

// SetDetail function to set map detail of ValidationError.
func (ve *ValidationError) SetDetail(key, value string) {
	if ve.detail == nil {
		ve.detail = make(map[string]string)
	}
	ve.detail[key] = value
}

// IsValidationError function convert error as ValidationError.
func IsValidationError(err error) (ValidationError, bool) {
	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		return validationErr, true
	}
	return ValidationError{}, false
}
