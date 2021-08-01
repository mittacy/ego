package model

import (
	"bytes"
	"html/template"
	"strings"
)

var modelTemplate = `
{{- /* delete empty line */ -}}
package model

type {{ .Name }} struct {
	Id int64
}

func (*{{ .Name }}) TableName() string {
	return "{{ .NameLower }}"
}

`

type Model struct {
	Name      string
	NameLower string
}

func (s *Model) execute() ([]byte, error) {
	s.NameLower = strings.ToLower(s.Name)
	s.Name = strings.Title(s.NameLower)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
