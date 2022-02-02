package job

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"text/template"
)

var taskTemplate = `
{{- /* delete empty line */ -}}
package job_payload

const {{ .Name }}TypeName = "{{ .NameLower }}:hello"

// Payload 任务数据
type {{ .Name }}Payload struct {
	RequestId string
}

`

type Task struct {
	AppName   string
	Name      string
	NameLower string
}

func (s *Task) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("model").Parse(taskTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
