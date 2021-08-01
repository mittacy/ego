package api

import (
	"bytes"
	"html/template"
	"strings"
)

var wireTemplate = `
func Init{{ .Name }}Api(db *gorm.DB, cache *redis.Pool) api.{{ .Name }} {
	wire.Build(data.New{{ .Name }}, service.New{{ .Name }}, api.New{{ .Name }})
	return api.{{ .Name }}{}
}
`

type wire struct {
	AppName string
	Name    string
}

func NewWire(appName, name string) wire {
	w := wire{
		AppName: appName,
		Name: name,
	}

	w.Name = strings.Title(strings.ToLower(w.Name))
	return w
}

func (s *wire) execute() ([]byte, error) {
	buf := new(bytes.Buffer)

	tmpl, err := template.New("wire").Parse(wireTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
