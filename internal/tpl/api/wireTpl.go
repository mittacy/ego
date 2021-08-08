package api

import (
	"bytes"
	"html/template"
	"strings"
)

var wireTemplate = `
func Init{{ .Name }}Api(db *gorm.DB, cache *redis.Pool) api.{{ .Name }} {
	customLogger := logger.NewCustomLogger("{{ .NameLower }}")
	{{ .NameLower }}Data := data.New{{ .Name }}(db, cache, customLogger)
	{{ .NameLower }}Service := service.New{{ .Name }}({{ .NameLower }}Data, customLogger)
	{{ .NameLower }}Api := api.New{{ .Name }}({{ .NameLower }}Service, customLogger)
	return {{ .NameLower }}Api
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
