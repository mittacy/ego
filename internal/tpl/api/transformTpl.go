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
	"github.com/mittacy/ego/library/gin/response"
	"github.com/mittacy/ego/library/log"
	"{{ .AppName }}/app/internal/model"
	"{{ .AppName }}/app/internal/validator/{{ .NameLower }}Validator"
)

var {{ .Name }} {{ .NameLower }}Transform

type {{ .NameLower }}Transform struct {}

// GetReply 详情响应
// @param data 数据库数据
func (ctl *{{ .NameLower }}Transform) GetReply(c *gin.Context, req interface{}, data model.{{ .Name }}) {
	reply{{ .Name }} := {{ .NameLower }}Validator.GetReply{}

	if err := copier.Copy(&reply{{ .Name }}, data); err != nil {
		log.New("{{ .NameLower }}").ErrorwWithTrace(c, "copier转化失败", "req", req, "data", data, "err", err)
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
