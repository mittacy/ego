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
	"github.com/gin-gonic/gin"
	"{{ .AppName }}/pkg/log"
)

var {{ .Name }} {{ .NameLower }}Api

func init() {
	l := log.New("{{ .NameLower }}")

	{{ .Name }} = {{ .NameLower }}Api{
		logger: l,
	}
}

type {{ .NameLower }}Api struct {
	logger *log.Logger
}

func (ctl *{{ .NameLower }}Api) Ping(c *gin.Context) {
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
