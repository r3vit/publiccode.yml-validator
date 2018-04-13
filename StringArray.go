package main

// StringArray can be a string or a slice
type StringArray []string

// UnmarshalYAML manages and resolves the duality string|array of go-yaml
// https://github.com/go-yaml/yaml/issues/100
func (a *StringArray) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		*a = []string{single}
	} else {
		*a = multi
	}
	return nil
}
