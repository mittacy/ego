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
	"github.com/mittacy/ego/library/gin/response"
	"github.com/mittacy/ego/library/log"
	"{{ .AppName }}/apierr"
	"{{ .AppName }}/app/internal/service"
	"{{ .AppName }}/app/internal/transform"
	"{{ .AppName }}/app/internal/validator/{{ .NameLower }}Validator"
)

var {{ .Name }} {{ .NameLower }}Api

type {{ .NameLower }}Api struct {}

func (ctl *{{ .NameLower }}Api) Get(c *gin.Context) {
	req := {{ .NameLower }}Validator.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}
	
	{{ .NameLower }}, err := service.{{ .Name }}.GetById(c, req.Id)
	if err != nil {
		response.CheckErrAndLog(c, log.New("{{ .NameLower }}"), req, "get{{ .Name }}", err, apierr.ResourceNoExist)
		return
	}

	transform.{{ .Name }}.GetReply(c, req, {{ .NameLower }})
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
