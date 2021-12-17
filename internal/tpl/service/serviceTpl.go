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
	"{{ .AppName }}/internal/data"
	"{{ .AppName }}/pkg/log"
)

// 一般情况下service应该只引用并控制自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型
// {{ .Name }} 服务说明注释
var {{ .Name }} {{ .NameLower }}Service

func init() {
	l := log.New("{{ .NameLower }}")

	{{ .Name }} = {{ .NameLower }}Service{
		logger: l,
		data: 	data.New{{ .Name }}(l),
	}
}

type {{ .NameLower }}Service struct {
	logger *log.Logger
	data   data.{{ .Name }}
}

func (ctl *{{ .NameLower }}Service) Ping() bool {
	return ctl.data.Ping()
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
