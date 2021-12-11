package api

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var transformTemplate = `
{{- /* delete empty line */ -}}
package transform

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"{{ .AppName }}/internal/model"
	"{{ .AppName }}/internal/validator/{{ .NameLower }}Validator"
	"{{ .AppName }}/pkg/log"
	"{{ .AppName }}/pkg/response"
)

var {{ .Name }} {{ .NameLower }}Transform

type {{ .NameLower }}Transform struct {
	logger *log.Logger
}

func init() {
	l := log.New("{{ .NameLower }}")

	{{ .Name }} = {{ .NameLower }}Transform{
		logger: l,
	}
}

// Get{{ .Name }}Reply 详情响应
// @param data 数据库数据
func (ctl *{{ .NameLower }}Transform) Get{{ .Name }}Reply(c *gin.Context, req interface{}, data *model.{{ .Name }}) {
	reply{{ .Name }} := {{ .NameLower }}Validator.GetReply{}

	if err := copier.Copy(&reply{{ .Name }}, data); err != nil {
		ctl.logger.CopierErrLog(err, req)
		response.Unknown(c)
		return
	}

	res := map[string]interface{}{
		"{{ .NameLower }}": reply{{ .Name }},
	}

	response.Success(c, res)
}
`

type Transform struct {
	Name      string
	NameLower string
	AppName   string
}

func (s *Transform) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("validator").Parse(transformTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
