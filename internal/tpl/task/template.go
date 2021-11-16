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
	"{{ .AppName }}/pkg/log"
	"github.com/robfig/cron/v3"
)

type {{ .Name }} struct {
	logger *log.Logger
}

func New{{ .Name }}(logger *log.Logger) *{{ .Name }} {
	return &{{ .Name }}{logger: logger}
}

func (t *{{ .Name }}) Name() string {
	return "{{ .NameLower }}Task"
}

func (t *{{ .Name }}) Spec() string {
	return "0 8 * * ?"
}

func (t *{{ .Name }}) Job() cron.Job {
	return t
}

func (t *{{ .Name }}) Run() {
	// do something
	t.logger.Info("Hello, this is the {{ .NameLower }} task")
}

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
