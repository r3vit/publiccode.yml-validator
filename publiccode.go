package main

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
