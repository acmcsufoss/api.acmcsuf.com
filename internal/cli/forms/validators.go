package forms

import (
	"errors"
)

// Used for fields that must contain a value
func ValidateNonEmpty() func(string) error {
	return func(field string) error {
		if field == "" {
			return errors.New("field must contain a value")
		}
		return nil
	}
}
