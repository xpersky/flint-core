package views

import (
	"bytes"
	"text/template"
)

func GetLong() string {

	buffer := &bytes.Buffer{}
	tmpl, err := template.ParseFiles("templates/long.tmpl")

	if err != nil {
		return buffer.String()
	}

	err = tmpl.Execute(buffer, nil)
	if err != nil {
		return buffer.String()
	}

	return buffer.String()
}

