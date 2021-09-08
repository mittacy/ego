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
	"{{ .AppName }}/app/model"
	"{{ .AppName }}/app/validator/{{ .NameLower }}Validator"
	"{{ .AppName }}/pkg/log"
	"{{ .AppName }}/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type {{ .Name }} struct {
	logger *log.Logger
}

func New{{ .Name }}(logger *log.Logger) {{ .Name }} {
	return {{ .Name }}{logger: logger}
}

// {{ .Name }}Pack 数据库数据转化为响应数据
// @param data 数据库数据
// @return reply 响应体数据
// @return err
func (ctl *{{ .Name }}) {{ .Name }}Pack(data *model.{{ .Name }}) (*{{ .NameLower }}Validator.GetReply, error) {
	reply := {{ .NameLower }}Validator.GetReply{}

	if err := copier.Copy(&reply, data); err != nil {
		return nil, err
	}

	return &reply, nil
}

// {{ .Name }}sPack 数据库数据转化为响应数据
// @param data 数据库数据
// @return reply 响应体数据
// @return err
func (ctl *{{ .Name }}) {{ .Name }}sPack(data []model.{{ .Name }}) (reply []{{ .NameLower }}Validator.ListReply, err error) {
	err = copier.Copy(&reply, &data)
	return
}

// GetReply 详情响应包装
// @param data 数据库数据
func (ctl *{{ .Name }}) GetReply(c *gin.Context, data *model.{{ .Name }}) {
	reply, err := ctl.{{ .Name }}Pack(data)
	if err != nil {
		response.CopierErrAndLog(c, ctl.logger, err)
		return
	}

	res := map[string]interface{}{
		"{{ .NameLower }}": reply,
	}

	response.Success(c, res)
}

// ListReply 列表响应包装
// @param data 数据库列表数据
// @param totalSize 记录总数(ctl *{{ .Name }}) 
func (ctl *{{ .Name }}) ListReply(c *gin.Context, data []model.{{ .Name }}, totalSize int64) {
	list, err := ctl.{{ .Name }}sPack(data)
	if err != nil {
		response.CopierErrAndLog(c, ctl.logger, err)
		return
	}

	res := map[string]interface{}{
		"list":       list,
		"total_size": totalSize,
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
