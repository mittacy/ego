package api

import (
	"bytes"
	"html/template"
	"strings"
)

var wireTemplate = `
func Init{{ .Name }}Api(db *gorm.DB, cache *redis.Pool) api.{{ .Name }} {
	customLogger := logger.NewCustomLogger("{{ .NameLower }}")
	i{{ .Name }}Service := data.New{{ .Name }}(db, cache, customLogger)
	apiI{{ .Name }}Service := service.New{{ .Name }}(i{{ .Name }}Service, customLogger)
	{{ .NameLower }} := api.New{{ .Name }}(apiI{{ .Name }}Service, customLogger)
	return {{ .NameLower }}
}
`

type wire struct {
	AppName   string
	Name      string
	NameLower string
}

func NewWire(appName, name string) wire {
	w := wire{
		AppName: appName,
		Name:    name,
	}

	w.Name = strings.Title(strings.ToLower(w.Name))
	return w
}

func (s *wire) execute() ([]byte, error) {
	s.NameLower = strings.ToLower(s.Name)
	s.Name = strings.Title(s.NameLower)

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
