package job

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"text/template"
)

var processorTemplate = `
{{- /* delete empty line */ -}}
package job_process

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego/library/log"
	"{{ .AppName }}/app/job/job_payload"
)

// Processor 任务处理器
type {{ .Name }}Processor struct {
	l *log.Logger
}

func New{{ .Name }}() *{{ .Name }}Processor {
	return &{{ .Name }}Processor{
		l: log.New("{{ .NameLower }}_job"),
	}
}

func (processor *{{ .Name }}Processor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p job_payload.{{ .Name }}Payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// call service
	// service.Biz.Do()
	processor.l.Info("do something")

	return nil
}

`

type Processor struct {
	AppName   string
	Name      string
	NameLower string
}

func (s *Processor) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("model").Parse(processorTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
