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
	"{{ .AppName }}/internal/service"
	"{{ .AppName }}/internal/transform"
	"{{ .AppName }}/pkg/log"
	"github.com/gin-gonic/gin"
)

type {{ .Name }}Api struct {
	service 	service.{{ .Name }}
	transform   transform.{{ .Name }}
	logger      *log.Logger
}

func New{{ .Name }}() {{ .Name }}Api {
	l := log.New("{{ .NameLower }}")
	return {{ .Name }}Api{
		logger:    l,
		service: service.New{{ .Name }}(l),
		transform: transform.New{{ .Name }}(l),
	}
}

func (ctl *{{ .Name }}Api) Ping(c *gin.Context) {
	c.JSON(200, "success")
}

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
