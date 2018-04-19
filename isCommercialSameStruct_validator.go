package main

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// isCommercialSameStruct return false if the Passed Field (param) is not "contract".
func isCommercialSameStruct(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	if fl.Parent().FieldByName(param).String() == "contract" {
		// field "param" is not empty
		if field.String() == "" {
			// current field is empty. Should not be empty!
			return false
		}

	}

	return true

}
