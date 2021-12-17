package model

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var modelTemplate = `
{{- /* delete empty line */ -}}
package model

type {{ .Name }} struct {
	Id int64
}

func (*{{ .Name }}) TableName() string {
	return "{{ .NameLower }}"
}

const (
	// 状态
	{{ .Name }}StateDeleted = 100	// 删除
)

`

type Model struct {
	Name      string
	NameLower string
}

func (s *Model) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
