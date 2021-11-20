package data

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var dataTemplate = `
package data

import (
	"fmt"
	"{{ .AppName }}/pkg/cache"
	"{{ .AppName }}/pkg/log"
	"{{ .AppName }}/pkg/mysql"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type {{ .Name }} struct {
	db 	   *gorm.DB
	cache  *redis.Client
	cacheKeyPre string
	logger *log.Logger
}

func New{{ .Name }}(logger *log.Logger) {{ .Name }} {
	return {{ .Name }}{
		logger: logger,
		db:    	mysql.NewClientByName("localhost"),
		cache: 	cache.NewClientByName("localhost", 0),
		cacheKeyPre: fmt.Sprintf("%s:{{ .NameLower }}", viper.GetString("APP_NAME")),
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
