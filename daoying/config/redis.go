package config

var (
	redisConfig = new(RedisConfig)
)

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"user_name"`
	PassWord string `mapstructure:"pass_word"`
}

func GetRedisConfig() *RedisConfig {
	err := GetConfig("redis", redisConfig)
	if err != nil {
		panic(err)
	}
	return redisConfig
}
