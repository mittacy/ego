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
	"{{ .AppName }}/pkg/log"
)

type {{ .Name }} struct {
	{{ .NameLower }}Data I{{ .Name }}Data
	logger *log.Logger
}

func New{{ .Name }}({{ .NameLower }}Data I{{ .Name }}Data, logger *log.Logger) {{ .Name }} {
	return {{ .Name }}{
		{{ .NameLower }}Data: {{ .NameLower }}Data,
		logger: logger,
	}
}

type I{{ .Name }}Data interface {
	PingData()
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
