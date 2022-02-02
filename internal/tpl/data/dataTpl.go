package data

import (
	"bytes"
	"github.com/mittacy/ego/internal/utils"
	"html/template"
)

var dataTemplate = `
{{- /* delete empty line */ -}}
package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego/library/mysql"
	"github.com/mittacy/ego/library/redis"
	"github.com/spf13/viper"
	"{{ .AppName }}/app/internal/model"
)

type {{ .Name }} struct {
	mysql.Gorm
	redis.GoRedis
	cacheKeyPre string
}

func New{{ .Name }}() {{ .Name }} {
	return {{ .Name }}{
		Gorm: mysql.Gorm{
			MysqlConfName: "localhost",
		},
		GoRedis: redis.GoRedis{
			RedisConfName: "localhost",
			RedisDB:       0,
		},
		cacheKeyPre: fmt.Sprintf("%s:{{ .NameLower }}", viper.GetString("APP_NAME")),
	}
}

func (ctl *{{ .Name }}) GetById(c *gin.Context, id int64) (model.{{ .Name }}, error) {
	return model.{{ .Name }}{Id: id}, nil
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
