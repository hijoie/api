package config

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	etcdAddr = "http://127.0.0.1:2379"
)

// func GetConfig(key string,value interface{}) (bizerr error)  {
// 	fmt.Println(key)
// 	mc := viper.New()
// 	bizerr = mc.AddRemoteProvider("etcd", etcdAddr, key)
// 	if bizerr != nil {
// 		return fmt.Errorf("add remote provider error %w",bizerr)
// 	}
// 	mc.SetConfigType("yaml")
// 	bizerr = mc.ReadRemoteConfig()
// 	if bizerr != nil {
// 		return fmt.Errorf("read remote config error %w",bizerr)
// 	}
// 	bizerr = mc.Unmarshal(value)
// 	if bizerr != nil {
// 		return fmt.Errorf("unmarshal config error %w",bizerr)
// 	}
// 	return
// }

func InitConfig(name string) {
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read in config error %w", err))
	}
}

func GetConfig(key string, value interface{}) (err error) {
	return viper.UnmarshalKey(key, value)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func SetEtcdAddr(addr string) {
	etcdAddr = addr
}
