package database

import (
	"fmt"
)

func (m *Manager) Run() (err error) {

	choice := -1

	var file string
	for choice != 0 {
		fmt.Println(">>>Choose a database<<<")
		fmt.Println(">1:sqlite3")
		fmt.Println(">2:mysl")
		fmt.Println(">0: exit")
		fmt.Print(">")
		_, err = fmt.Scan(&choice)

		if err != nil {
			choice = -1
			fmt.Println("> this is not an usable value")
		}

		switch choice {
		case 0:
			fmt.Println("Exit....")
		case 1:
			if file == "" {
				fmt.Print("> Enter the name of the config file : ")
				fmt.Scan(&file)
				choice = -1
			}

			err = m.InitSqliteDB(file)
			if err != nil {
				choice = -1
				fmt.Println("> erreur : ", err)
			}

			err = m.TestDB()
			if err != nil {
				choice = -1
				fmt.Println("> erreur : ", err)
			}
			err = m.Exec()
			if err != nil {
				fmt.Println("> erreur : ", err)
				choice = -1
			}
			m.Close()
		case 2:
			var conf Configuration
			err = m.InitMysqlDB(conf)
			if err != nil {
				choice = -1
				fmt.Println("> erreur : ", err)
			}
			err = m.TestDB()

			if err != nil {
				fmt.Println("> erreur : ", err)
				choice = -1
			}
			err = m.Exec()
			if err != nil {
				fmt.Println("> erreur : ", err)
				choice = -1
			}
			m.Close()
		default:
			fmt.Println("command does'nt exist...")
			choice = -1
		}

	}
	return err
}

func (m Manager) Exec() (err error) {
	//create a menu in the console
	choice := -1

	for choice != 0 {
		fmt.Println(">>>Choose an action<<<")
		fmt.Println(">1:create a table")
		fmt.Println(">2:show all tables")
		fmt.Println(">0: exit")
		fmt.Print(">")
		_, err = fmt.Scan(&choice)

		if err != nil {
			choice = -1
			fmt.Println("> this is not an usable value")
		}

		switch choice {
		case 0:
			fmt.Println("Exit....")
		case 1:

			err = m.CreateTable(User{})
			if err != nil {
				choice = -1
				fmt.Println("> erreur : ", err)
			}
			var user User
			fmt.Println("Username : ")
			fmt.Scan(&user.Name)
			fmt.Println("Password : ")
			fmt.Scan(&user.Password)
			err = m.CreateUser(user)
			if err != nil {
				choice = -1
				fmt.Println("> erreur : ", err)
			}

		case 2:
			err = m.ShowTables()
			if err != nil {
				choice = -1
				fmt.Println("> erreur : ", err)
			}
		}
	}

	return err

}
