package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func DNS() string {
	dbName := viper.GetString("mysql.dbname")
	userName := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	path := viper.GetString("mysql.path")
	port := viper.GetString("mysql.port")
	config := viper.GetString("mysql.config")
	return userName + ":" + password + "@tcp(" + path + ":" + port + ")/" + dbName + "?" + config
}

func InitMysql() {
	// 自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)
	//db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{
	db, err := gorm.Open(mysql.Open(DNS()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("err: ", err)
	}
	DB = db
	fmt.Println("config mysql inited ....")
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
}
