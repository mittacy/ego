package api

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var wireTemplate = `
func Init{{ .Name }}Api(db *gorm.DB, cache *redis.Pool) api.{{ .Name }} {
	customLogger := logger.NewCustomLogger("{{ .NameLower }}")
	{{ .NameLower }}Data := data.New{{ .Name }}(db, cache, customLogger)
	{{ .NameLower }}Service := service.New{{ .Name }}({{ .NameLower }}Data, customLogger)
	return api.New{{ .Name }}({{ .NameLower }}Service, customLogger)
}
`

var wireVarTemplate = `
  {{ .NameLower }}Api api.{{ .Name }}
`

var wireVarInitTemplate = `
  {{ .NameLower }}Api = Init{{ .Name }}Api(db.ConnectGorm("MYSQLKEY"), cache.ConnRedis("REDISKEY"))
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

	w.Name = utils.StringFirstUpper(w.Name)
	return w
}

func (s *wire) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

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

func (s *wire) executeVar() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("wire").Parse(wireVarTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *wire) executeVarInit() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("wire").Parse(wireVarInitTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
