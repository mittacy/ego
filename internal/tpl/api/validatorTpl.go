package api

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
	"strings"
)

var validatorTemplate = `
{{- /* delete empty line */ -}}
package {{ .NameLower }}Validator

type GetReq struct {
	Id int64 ${backquote}form:"id" json:"id" binding:"required,min=1"${backquote}
}
type GetReply struct {
	Id int64 ${backquote}json:"id"${backquote}
}
`

type Validator struct {
	Name      string
	NameLower string
}

func (s *Validator) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	validatorTemplate = strings.Replace(validatorTemplate, "${backquote}", "`", -1)
	tmpl, err := template.New("validator").Parse(validatorTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
