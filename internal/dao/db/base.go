package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall/internal/core"
	"mall/internal/logger"
	"sync"
	"time"
)

var dbInstance map[string]*gorm.DB

var once sync.Once

func initMysql() {
	mysqlConfig := core.GlobalConfig.Mysql

	tmpInstance := make(map[string]*gorm.DB, 0)
	// 数据库配置是一个数组
	for _, conf := range mysqlConfig {
		config := logger.GormLoggerConfig{
			SlowThreshold: conf.SlowThreshold * time.Millisecond,
			TraceLog:      conf.TraceLog,
		}
		newLogger := logger.NewGormLog(config)

		gormConfig := gorm.Config{}
		gormConfig.Logger = newLogger

		db, err := gorm.Open(mysql.Open(conf.Dsn), &gormConfig)
		if err != nil {
			fmt.Println(err)
		}

		tmpInstance[conf.Instance] = db
	}

	dbInstance = tmpInstance

}

// 单例模式：获取数据库实例
func GetDbInstance(db string) *gorm.DB {
	if db == "" {
		db = "default"
	}
	if dbInstance != nil {
		return dbInstance[db]
	}

	once.Do(func() {
		initMysql()
	})

	return dbInstance[db]
}
