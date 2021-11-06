package data

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var dataTemplate = `
package data

import (
	"{{ .AppName }}/pkg/cache"
	"{{ .AppName }}/pkg/log"
	"{{ .AppName }}/pkg/mysql"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type {{ .Name }} struct {
	db 	   *gorm.DB
	cache  *redis.Client
	logger *log.Logger
}

func New{{ .Name }}(logger *log.Logger) {{ .Name }} {
	return {{ .Name }}{
		db:    	mysql.NewClientByName("localhost"),
		cache: 	cache.NewClientByName("localhost", 0),
		logger: logger,
	}
}

func (ctl *{{ .Name }}) PingData() {}

`

type Data struct {
	AppName   string
	Name      string
	NameLower string
}

func (s *Data) execute() ([]byte, error) {
	s.Name = utils.StringFirstUpper(s.Name)
	s.NameLower = utils.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("data").Parse(dataTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
