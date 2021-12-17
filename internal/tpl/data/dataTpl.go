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
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"{{ .AppName }}/pkg/cache"
	"{{ .AppName }}/pkg/log"
	"{{ .AppName }}/pkg/mysql"
)

type {{ .Name }} struct {
	db          *gorm.DB
	cache       *redis.Client
	cacheKeyPre string
	logger      *log.Logger
}

func New{{ .Name }}(l *log.Logger) {{ .Name }} {
	return {{ .Name }}{
		db:          mysql.NewClientByName("localhost"),
		cache:       cache.NewClientByName("localhost", 0),
		cacheKeyPre: fmt.Sprintf("%s:{{ .NameLower }}", viper.GetString("APP_NAME")),
		logger:      l,
	}
}

func (ctl *{{ .Name }}) Ping() bool {
	return true
}

/*
 * 以下为查询缓存KEY方法
 */
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
