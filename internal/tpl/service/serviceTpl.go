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
	"{{ .AppName }}/app/api"
)

type {{ .Name }} struct {
	{{ .NameLower }}Data I{{ .Name }}Data
}

// 编写实现api层中的各个service接口的构建方法

func New{{ .Name }}({{ .NameLower }}Data I{{ .Name }}Data) api.I{{ .Name }}Service {
	return &{{ .Name }}{ {{ .NameLower }}Data: {{ .NameLower }}Data }
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
