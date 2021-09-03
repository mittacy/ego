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

type CreateReq struct {}
type CreateReply struct{}

type DeleteReq struct {}
type DeleteReply struct{}

type UpdateReq struct {
	UpdateType int ${backquote}json:"update_type" binding:"required,oneof=1"${backquote}
}

type UpdateInfoReq struct {}
type UpdateInfoReply struct{}

type GetReq struct {}
type GetReply struct {}

type ListReq struct {}
type ListReply struct {}

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
