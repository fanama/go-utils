package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (m *Manager) InitMysqlDB(conf Configuration) (err error) {

	conf, err = ReadConfig("database")

	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})

	return err
}

func (m *Manager) InitSqliteDB(name string) (err error) {

	m.db, err = gorm.Open(sqlite.Open(name+".db"), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})

	return err
}

func (m Manager) Close() {
	sqlDB, err := m.db.DB()

	if err != nil {
		log.Fatal(err)
	}

	sqlDB.Close()
}
