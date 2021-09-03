package api

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var apiTemplate = `
{{- /* delete empty line */ -}}
package api

import (
	"{{ .AppName }}/app/service"
	"{{ .AppName }}/app/transform"
	"{{ .AppName }}/pkg/logger"
	"github.com/gin-gonic/gin"
)

type {{ .Name }} struct {
	{{ .NameLower }}Service service.{{ .Name }}
	transform   transform.{{ .Name }}
	logger      *logger.CustomLogger
}

func New{{ .Name }}({{ .NameLower }}Service service.{{ .Name }}, logger *logger.CustomLogger) {{ .Name }} {
	return {{ .Name }}{
		{{ .NameLower }}Service: {{ .NameLower }}Service,
		transform: transform.New{{ .Name }}(logger),
		logger:    logger,
	}
}

func (ctl *{{ .Name }}) Ping(c *gin.Context) {}

`

type Api struct {
	AppName   string
	Name      string
	NameLower string
}

func (s *Api) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

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
