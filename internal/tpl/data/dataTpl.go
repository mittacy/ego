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
	{{- if .InjectMongo}}
	"github.com/mittacy/ego/library/eMongo"
	{{- end}}
	{{- if .InjectMysql}}
	"github.com/mittacy/ego/library/eMysql"
	{{- end}}
	{{- if .InjectRedis}}
	"github.com/mittacy/ego/library/eRedis"
	{{- end}}
	"github.com/spf13/viper"
	"{{ .AppName }}/app/internal/model"
)

type {{ .Name }} struct {
	{{- if .InjectMysql}}
	eMysql.EGorm
	{{- end}}
	{{- if .InjectMongo}}
	eMongo.EMongo
	{{- end}}
	{{- if .InjectRedis}}
	eRedis.ERedis
	cacheKeyPre string
	{{- end}}
	{{- if .InjectHttp}}
	host string
	{{- end}}
}

func New{{ .Name }}() {{ .Name }} {
	return {{ .Name }}{
		{{- if .InjectMysql}}
		EGorm: eMysql.EGorm{
			MysqlConfName: "localhost",
		},
		{{- end}}
		{{- if .InjectMongo}}
		EMongo: eMongo.EMongo{
			MongoConfName: "localhost",
			CollationName: "user",
		},
		{{- end}}
		{{- if .InjectRedis}}
		ERedis: eRedis.ERedis{
			RedisConfName: "localhost",
			RedisDB:       0,
		},
		cacheKeyPre: fmt.Sprintf("%s:user", viper.GetString("APP_NAME")),
		{{- end}}
		{{- if .InjectHttp}}
		host: viper.GetString("{{ .Name }}"),
		{{- end}}
	}
}

func (ctl *{{ .Name }}) GetById(c *gin.Context, id int64) (model.{{ .Name }}, error) {
	return model.{{ .Name }}{Id: id}, nil
}

{{if .InjectRedis}}
/*
 * 以下为查询缓存KEY方法
 */
{{- end}}

{{if .InjectHttp}}
/*
 * 以下为服务接口Uri路由
 */
func (ctl *{{ .Name }}) AddUri() string {
	return ""
}
{{- end}}
`

type Data struct {
	AppName     string
	Name        string
	NameLower   string
	InjectMysql bool
	InjectMongo bool
	InjectRedis bool
	InjectHttp  bool
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

func (s *Data) parseInject(databaseHandle []string) {
	for _, v := range databaseHandle {
		if v == InjectMysql {
			s.InjectMysql = true
		} else if v == InjectMongo {
			s.InjectMongo = true
		} else if v == InjectRedis {
			s.InjectRedis = true
		} else if v == InjectHttp {
			s.InjectHttp = true
		}
	}
}
