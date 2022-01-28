package mysql

const (
	dbDSNFormat = "%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local"
)

type Conf struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	Params   string
}

var connectConf map[string]Conf // 连接配置池

// InitMysqlConf 函数功能说明
// Example:
// c = map[string]Conf{
//		"localhost": {
//			Host:     viper.GetString("DB_CORE_RW_HOST"),
//			Port:     viper.GetInt("DB_CORE_RW_PORT"),
//			Database: viper.GetString("DB_CORE_RW_DATABASE"),
//			User:     viper.GetString("DB_CORE_RW_USERNAME"),
//			Password: viper.GetString("DB_CORE_RW_PASSWORD"),
//		},
//	}
// @param c
func InitMysqlConf(c map[string]Conf) {
	connectConf = c
}
