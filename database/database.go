package database

import (
	"fmt"
	"strings"
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

func (m Manager) CreateTable() (err error) {

	// ask for table name and column names and types
	// ask for number of columns
	// ask for primary key

	fmt.Println("Creating table------------------------------------------------------")
	fmt.Println("Enter table name")
	var tmp string
	fmt.Scan(&tmp)
	var tableName string
	fmt.Scanln(&tableName)
	fmt.Println("Enter number of columns")
	var numCols int
	fmt.Scanln(&numCols)
	fmt.Println("Enter primary key")
	var primaryKey string
	fmt.Scanln(&primaryKey)
	fmt.Println("Enter column names and types")
	var cols []string
	for i := 0; i < numCols; i++ {
		fmt.Println(i, ": Enter column name")
		var colName string
		fmt.Scanln(&colName)
		fmt.Println(i, ": Enter column type")
		var colType string
		fmt.Scanln(&colType)
		cols = append(cols, colName+" "+colType)
	}

	// create table
	// create a list of column names and types

	request := "CREATE TABLE " + tableName + " (" + primaryKey + " INTEGER PRIMARY KEY AUTOINCREMENT, " + strings.Join(cols, ",") + ")"

	fmt.Println(request)

	// err = m.db.Raw(request).Error

	return err
}

func (m Manager) ShowTables() (err error) {
	fmt.Println("Showing tables------------------------------------------------------")
	var tables []string
	err = m.db.Table("sqlite_master").Select("name").Where("type = ?", "table").Pluck("name", &tables).Error
	if err != nil {
		return err
	}
	fmt.Println(tables)
	return err
}
