package job

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"text/template"
)

var processorTemplate = `
{{- /* delete empty line */ -}}
package {{ .NameLower }}JobProcess

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"{{ .AppName }}/pkg/async"
	"{{ .AppName }}/pkg/log"
	"{{ .AppName }}/interface/job/{{ .NameLower }}Job/{{ .NameLower }}JobTask"
)

func NewProcessor() *Processor {
	return &Processor{
		l: async.GetLogger(),
	}
}

// Processor 任务处理器, 实现 asynq.Handler 接口
type Processor struct {
	l *log.Logger
}

func (processor *Processor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p {{ .NameLower }}JobTask.Payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	
	// do work...
	processor.l.Infof("数据: %+v", p)

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
