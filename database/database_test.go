package database

import "testing"

// go test -v ./database -run TestDatabaseSqlite
func TestDatabaseSqlite(t *testing.T) {
	m := Manager{}

	m.InitSqliteDB("test")

	type User struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	m.db.AutoMigrate(&User{})

	m.db.Create(&User{Name: "test", Age: 10})

	var user User

	m.db.Find(&User{}).Scan(&user)

	if user.Name != "test" {
		t.Errorf("User name should be test, but got %s", user.Name)
	}

	if user.Age != 10 {
		t.Errorf("User age should be 10, but got %d", user.Age)
	}
	t.Logf("User name is %s, age is %d", user.Name, user.Age)

	m.db.Delete(&user)

}

func TestListTableSlite(t *testing.T) {
	m := Manager{}

	m.InitSqliteDB("test")

	if m.ShowTables() != nil {
		t.Errorf("ShowTables should be nil, but got %v", m.ShowTables())
	}

}
