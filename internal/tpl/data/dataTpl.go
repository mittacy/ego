package data

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var dataTemplate = `
package data

import (
	"{{ .AppName }}/app/service"
	"{{ .AppName }}/pkg/log"
	"{{ .AppName }}/pkg/store/cache"
	"gorm.io/gorm"
)

// 实现service层中的data接口

type {{ .Name }} struct {
	db 	   *gorm.DB
	redis  *cache.Redis
	logger *log.Logger
}

func New{{ .Name }}(db *gorm.DB, redis *cache.Redis, logger *log.Logger) service.I{{ .Name }}Data {
	return &{{ .Name }}{
		db:    	db,
		redis: 	redis,
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
