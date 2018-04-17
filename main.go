package main

import (
	"fmt"
	"io/ioutil"
	"log"

	validator "gopkg.in/go-playground/validator.v9"
	yaml "gopkg.in/yaml.v2"
)

// PublicCode describe a publiccode.yml file with struct tags for validation.
type PublicCode struct {
	Version     string      `validate:"required"`
	URL         string      `validate:"required,url"`
	UpstreamURL StringArray `yaml:"upstream-url" validate:"dive,url"`
	Legal       struct {
		MainCopyrightOwner string `yaml:"main-copyright-owner"`
		RepoOwner          string `yaml:"repo-owner" validate:"required"`
		AuthorsFile        string `yaml:"authors-file"`
		License            string `validate:"required,spdx"`
	}
	Maintenance struct {
		Type              string      `validate:"required,maintenance"`
		Until             string      `validate:"required,date"`
		Maintainer        StringArray `validate:"required"`
		TechnicalContacts []struct {
			Name        string `validate:"required"`
			Email       string `validate:"isnotemptyss=Name,email"`
			Affiliation string `validate:"isnotemptyss=Name"`
		} `yaml:"technical-contacts" validate:"dive"`
	}
	Description []struct {
		Name      string `validate:"required"`
		Logo      StringArray
		Shortdesc []struct {
			En string
		}
		Longdesc []struct {
			En string
		}
		Screenshots []string
		Videos      []string
		Version     string
		Released    string `validate:"isnotemptyss=Version,date"`
		Platforms   StringArray
	} `validate:"dive"`
}

// Use a single instance of Validate, it will cache struct info.
var validate *validator.Validate

func main() {
	var pc PublicCode

	// Read data from file.
	data, err := ioutil.ReadFile("publiccode.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the yaml into Publiccode struct.
	err = yaml.UnmarshalStrict(data, &pc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Generate a new validator and add new validation rules.
	validate = validator.New()
	err = validate.RegisterValidation("spdx", spdxValidator)
	checkErr(err)
	err = validate.RegisterValidation("maintenance", maintenanceValidator)
	checkErr(err)
	err = validate.RegisterValidation("date", dateValidator)
	checkErr(err)
	err = validate.RegisterValidation("isnotemptyss", isNotEmptySameStruct)
	checkErr(err)
	// Apply validate rules on PublicCode struct.
	err = validate.Struct(pc)
	if err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Println("publiccode.yml is valid!")
	}

}

// checkErr for better readability. Gometalinter checks every validate.RegisterValidation err.
func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
