package task

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var modelTemplate = `
{{- /* delete empty line */ -}}
package task

import (
	"github.com/mittacy/ego/library/log"
	"github.com/robfig/cron/v3"
)

type {{ .Name }} struct {
	logger *log.Logger
}

func New{{ .Name }}(logger *log.Logger) *{{ .Name }} {
	return &{{ .Name }}{logger: logger}
}

func (t *{{ .Name }}) Name() string {
	return "{{ .NameLower }}"
}

func (t *{{ .Name }}) Spec() string {
	return "0 8 * * ?"
}

func (t *{{ .Name }}) Job() cron.Job {
	return t
}

func (t *{{ .Name }}) Run() {
	// do something
	t.logger.Infow("Hello, this is the example task", "task", t.Name())
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

	tmpl, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
