package main

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// Validate maintenance type.
// There are three types of maintenance: internal, contract, community or none:
// 	internal: maintenance is performed by the authors
// 	contract: maintenance is performed by a contracted party
// 	community: maintenance is done by the community
// 	none: this project is unmaintained, if you file a pull request you'll likely get no response
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
	"internal",
	"contract",
	"community",
	"none",
}
