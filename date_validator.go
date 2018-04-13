package main

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

// Validate date in the form (YYYY-MM-DD).
// All dates in publiccode.yml must follow the format "YYYY-MM-DD",
// which is one of the ISO8601 allowed encoding.
// This is the only allowed encoding though, so not the full ISO8601
// is allowed for the date keys.*/
func dateValidator(fl validator.FieldLevel) bool {
	date := fl.Field().String()

	// Use the go time.Parse to check.
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return false
	}

	return true

}
