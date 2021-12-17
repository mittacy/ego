package service

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var serviceTemplate = `
{{- /* delete empty line */ -}}
package service

import (
	"{{ .AppName }}/internal/data"
	"{{ .AppName }}/pkg/log"
)

var {{ .Name }} {{ .NameLower }}Service

func init() {
	l := log.New("{{ .NameLower }}")

	{{ .Name }} = {{ .NameLower }}Service{
		logger: l,
		data: 	data.New{{ .Name }}(l),
	}
}

type {{ .NameLower }}Service struct {
	logger *log.Logger
	data   data.{{ .Name }}
}

func (ctl *{{ .NameLower }}Service) Ping() bool {
	return ctl.data.Ping()
}

`

type Service struct {
	AppName   string
	Name      string
	NameLower string
}

func (s *Service) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
