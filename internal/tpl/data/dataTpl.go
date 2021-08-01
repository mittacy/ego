package data

import (
	"bytes"
	"html/template"
	"strings"
)

var dataTemplate = `
package data

import (
	"{{ .AppName }}/app/service"
	"{{ .AppName }}/pkg/store/cache"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type {{ .Name }} struct {
	db *gorm.DB
	cache cache.CustomRedis
}

// 编写实现service层中的各个data接口的构建方法

func New{{ .Name }}(db *gorm.DB, cacheConn *redis.Pool) service.I{{ .Name }}Data {
	r := cache.GetCustomRedisByPool(cacheConn, "{{ .NameLower }}")

	return &{{ .Name }}{
		db:    db,
		cache: r,
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
	s.NameLower = strings.ToLower(s.Name)
	s.Name = strings.Title(s.NameLower)

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
