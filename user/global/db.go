package global

import (
	"api/user/config"
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	rds           *redis.Client
	mydb          *gorm.DB
	cacheInstance *cache.Cache
	localCache    cache.LocalCache
)

func InitDB() {
	rds = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.GetRedisConfig().Host, config.GetRedisConfig().Port),
	})
	logrus.Info(rds.Ping(context.TODO()).Result())

	mydb = newMysql()

	cacheInstance = cache.New(&cache.Options{
		Redis: rds,
	})
	localCache = cache.NewTinyLFU(10000, time.Hour)
}

func newMysql() *gorm.DB {
	mc := config.GetMysqlConfig("user")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mc.UserName, mc.PassWord, mc.Host, mc.Port, mc.DbName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func GetRds() *redis.Client {
	return rds
}

func GetMyDB() *gorm.DB {
	return mydb
}

func GetCache() *cache.Cache {
	return cacheInstance
}

func GetLocalCache() cache.LocalCache {
	return localCache
}
