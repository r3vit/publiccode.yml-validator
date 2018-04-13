package main

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// isNotEmpty return false if the Passed Field (param) is not empty.
// It's used to manage conditional values.
// Eg: this field is valid IF field "param" is not empty.
//     			Name        string
//          Email       string `validate:"isnotemptyss=Name,email"`
// The param field Name must not be empty in order to have an Email and email is required.
func isNotEmptySameStruct(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	if fl.Parent().FieldByName(param).String() != "" {
		// field "param" is not empty
		if field.String() == "" {
			// current field is empty. Should not be empty!
			return false
		}

	}

	return true

}
