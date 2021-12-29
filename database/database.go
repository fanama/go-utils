package database

import (
	"fmt"
)

func (m *Manager) TestDB() (err error) {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()

	if err != nil {
		return err
	}

	fmt.Println("database created successfully")
	return err
}

func (m *Manager) CreateTable(typeObject interface{}) (err error) {
	fmt.Println("Creating table------------------------------------------------------")
	err = m.db.AutoMigrate(typeObject)

	return err
}
func (m *Manager) CreateUser(user User) (err error) {
	fmt.Println("creating element----------------------------")
	err = m.db.Create(&user).Error

	return err
}

func (m Manager) ShowTables() (err error) {
	// gorm show config database in terminal
	fmt.Println("show tables------------------------------------------------------")
	var tables []string
	err = m.db.Raw("SHOW TABLES").Scan(&tables).Error
	fmt.Println("tables : ", tables)

	if err != nil {
		err = m.db.Table("sqlite_schema").Select("name").Scan(&tables).Error
	}

	if err != nil {
		fmt.Println("error : ", err)
	}

	return err
}
