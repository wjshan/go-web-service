package units

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBInstance = &DBMaker{}

type DBMaker struct {
	pgsqlInstance *gorm.DB
}

func (i *DBMaker) InitPostgresDb(dsn string) error {
	dbInstance, ok := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	i.pgsqlInstance = dbInstance
	log := LogInstance.GetLog("DBInit")
	if ok != nil {
		log.Fatalf("数据库链接错误:%s", ok.Error())
		return ok
	}
	sqlDB, _ := dbInstance.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	return nil
}
func (i *DBMaker) GetPGSQL() *gorm.DB {
	return i.pgsqlInstance
}
