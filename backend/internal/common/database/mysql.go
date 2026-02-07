package database

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitMySQL 从 viper 配置中初始化 MySQL 数据库并自动迁移模型
// configPrefix 指定在 viper 中的配置前缀，例如 "services.logistics.mysql"
func InitMySQL(configPrefix string, models ...interface{}) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString(configPrefix+".user"),
		viper.GetString(configPrefix+".password"),
		viper.GetString(configPrefix+".host"),
		viper.GetString(configPrefix+".port"),
		viper.GetString(configPrefix+".database"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to MySQL database at %s: %v\n", configPrefix, err)
		return nil, err
	}

	if len(models) > 0 {
		err = db.AutoMigrate(models...)
		if err != nil {
			fmt.Printf("Failed to auto-migrate for %s: %v\n", configPrefix, err)
			return nil, err
		}
	}

	return db, nil
}

// MustInitMySQL 初始化 MySQL，如果失败则直接退出程序
func MustInitMySQL(configPrefix string, models ...interface{}) *gorm.DB {
	db, err := InitMySQL(configPrefix, models...)
	if err != nil {
		os.Exit(1)
	}
	return db
}
