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

type {{ .Name }} struct {
	data data.{{ .Name }}
	logger *log.Logger
}

func New{{ .Name }}(logger *log.Logger) {{ .Name }} {
	return {{ .Name }}{
		data: data.New{{ .Name }}(logger),
		logger: logger,
	}
}

func (ctl *{{ .Name }}) Ping() {}

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
