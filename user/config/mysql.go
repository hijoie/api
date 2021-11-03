package config

const (
	//mysqlConfigPath = "/configs/server/mysql"
)

var (
	mysqlConfig = new(MysqlConfig)
)

type MysqlConfig struct {
	Host     string `json:"host,omitempty" mapstructure:"host"`
	Port     int    `json:"port,omitempty" mapstructure:"port"`
	DbName   string `json:"db_name,omitempty" mapstructure:"db_name"`
	UserName string `json:"user_name,omitempty" mapstructure:"user_name"`
	PassWord string `json:"pass_word,omitempty" mapstructure:"pass_word"`
}

func GetMysqlConfig() *MysqlConfig {
	err := GetConfig("mysql",mysqlConfig)
	if err != nil {
		panic(err)
	}
	return mysqlConfig
}