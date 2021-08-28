package service

import (
	"bytes"
	"html/template"
	"strings"
)

var serviceTemplate = `
{{- /* delete empty line */ -}}
package service

import (
	"{{ .AppName }}/pkg/logger"
)

type {{ .Name }} struct {
	{{ .NameLower }}Data I{{ .Name }}Data
	logger *logger.CustomLogger
}

func New{{ .Name }}({{ .NameLower }}Data I{{ .Name }}Data, logger *logger.CustomLogger) {{ .Name }} {
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
	s.NameLower = strings.ToLower(s.Name)
	s.Name = strings.Title(s.NameLower)

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
