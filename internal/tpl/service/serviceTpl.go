package service

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var serviceTemplate = `
{{- /* delete empty line */ -}}
package service

import (
	"github.com/gin-gonic/gin"
	"{{ .AppName }}/app/internal/data"
	"{{ .AppName }}/app/internal/model"
)

// 一般情况下service应该只引用并控制自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型
// {{ .Name }} 服务说明注释
var {{ .Name }} = {{ .NameLower }}Service{
	data: data.New{{ .Name }}(),
}

type {{ .NameLower }}Service struct {
	data   data.{{ .Name }}
}

func (ctl *{{ .NameLower }}Service) GetById(c *gin.Context, id int64) (model.{{ .Name }}, error) {
	return ctl.data.GetById(c, id)
}
`

type Service struct {
	AppName   string
	Name      string
	NameLower string
}

func (s *Service) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

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
