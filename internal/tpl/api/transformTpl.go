package api

import (
	"bytes"
	"html/template"
	"strings"
)

var transformTemplate = `
{{- /* delete empty line */ -}}
package {{ .NameLower }}Transform

import (
	"{{ .AppName }}/app/model"
	"{{ .AppName }}/app/validator/{{ .NameLower }}Validator"
	"{{ .AppName }}/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// {{ .Name }}Pack 数据库数据转化为响应数据
// @param data 数据库数据
// @return reply 响应体数据
// @return err
func {{ .Name }}Pack(data *model.{{ .Name }}) (*{{ .NameLower }}Validator.GetReply, error) {
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
func {{ .Name }}sPack(data []model.{{ .Name }}) (reply []{{ .NameLower }}Validator.ListReply, err error) {
	err = copier.Copy(&reply, &data)
	return
}

// {{ .Name }}ToReply 详情响应包装
// @param data 数据库数据
func {{ .Name }}ToReply(c *gin.Context, data *model.{{ .Name }}) {
	reply, err := {{ .Name }}Pack(data)
	if err != nil {
		response.CopierErrAndLog(c, err)
		return
	}

	res := map[string]interface{}{
		"{{ .NameLower }}": reply,
	}

	response.Success(c, res)
}

// {{ .Name }}sToReply 列表响应包装
// @param data 数据库列表数据
// @param totalSize 记录总数
func {{ .Name }}sToReply(c *gin.Context, data []model.{{ .Name }}, totalSize int64) {
	list, err := {{ .Name }}sPack(data)
	if err != nil {
		response.CopierErrAndLog(c, err)
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
	s.NameLower = strings.ToLower(s.Name)
	s.Name = strings.Title(s.NameLower)

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
