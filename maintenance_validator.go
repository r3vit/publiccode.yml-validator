package main

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// Validate maintenance type.
// There are three types of maintenance: commercial, community or none.
func maintenanceValidator(fl validator.FieldLevel) bool {
	maintenance := fl.Field().String()

	for _, maintenanceType := range maintenanceTypes {
		if maintenanceType == maintenance {
			return true
		}
	}
	return false
}

var maintenanceTypes = []string{
	"commercial",
	"community",
	"none",
}
