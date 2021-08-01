package api

import (
	"bytes"
	"html/template"
	"strings"
)

var apiTemplate = `
{{- /* delete empty line */ -}}
package api

import (
	"github.com/gin-gonic/gin"
)

type {{ .Name }} struct {
	{{ .NameLower }}Service I{{ .Name }}Service
}

func New{{ .Name }}({{ .NameLower }}Service I{{ .Name }}Service) {{ .Name }} {
	return {{ .Name }}{ {{ .NameLower }}Service: {{ .NameLower }}Service }
}

type I{{ .Name }}Service interface {
	Ping()
}

func (ctl *{{ .Name }}) Ping(c *gin.Context) {}

`

type Api struct {
	Name      string
	NameLower string
}

func (s *Api) execute() ([]byte, error) {
	s.NameLower = strings.ToLower(s.Name)
	s.Name = strings.Title(s.NameLower)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("api").Parse(apiTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
