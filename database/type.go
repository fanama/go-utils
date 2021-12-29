package database

import "gorm.io/gorm"

type Manager struct {
	db *gorm.DB
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
