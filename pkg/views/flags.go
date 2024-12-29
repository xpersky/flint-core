package views

import (
	"bytes"
	"text/template"
)

func GetConfigFlag() string {

	buffer := &bytes.Buffer{}
	tmpl, err := template.ParseFiles("templates/flags/config.tmpl")

	if err != nil {
		return buffer.String()
	}

	err = tmpl.Execute(buffer, nil)
	if err != nil {
		return buffer.String()
	}

	return buffer.String()
}

