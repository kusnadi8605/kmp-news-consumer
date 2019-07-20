package parser

import (
	"io/ioutil"

	yaml "github.com/go-yaml/yaml"
)

// LoadYAML load yaml format configuration
func LoadYAML(filename *string, v interface{}) error {
	raw, err := ioutil.ReadFile(*filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, v)
	if err != nil {
		return err
	}
	return nil
}

// LoadJSON load json format configuration
/*
func LoadJSON(filename *string, v interface{}) error {
	raw, err := ioutil.ReadFile(*filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, v)
	if err != nil {
		return err
	}

	return nil
}
*/
