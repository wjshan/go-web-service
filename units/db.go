package units

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitPostgresDb(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log := GetLog("DBInit")
	if err != nil {
		log.Errorf("数据库链接错误:%s", err.Error())
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	return nil
}

func GetDB() *gorm.DB {
	return db
}
