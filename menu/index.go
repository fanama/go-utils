package menu

import (
	"fmt"
	"log"

	"github.com/fanama/go-utils/fichier"
	"github.com/fanama/go-utils/mail"
)

func Display() {
	choice := -1
	for choice != 0 {
		fmt.Println("-------------------------------------------------------")
		fmt.Println("###Choose an action :")
		fmt.Println("- Create a json file: 1")
		fmt.Println("- Read a json file : 2")
		fmt.Println("- Send an Email : 3")
		fmt.Println("- Exit : 0")
		fmt.Print("Choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			fmt.Println("Exit....")
		case 1:
			var fileName string
			fmt.Print("fileName : ")
			fmt.Scan(&fileName)
			fichier.CreateJSON(fileName)
		case 2:
			var fileName string
			fmt.Print("fileName : ")
			fmt.Scan(&fileName)
			res, _ := fichier.ReadJSON(fileName)
			fmt.Println("resultat : ", res)
		case 3:
			m := mail.Configuration{}
			err := m.Run()
			if err != nil {
				log.Fatal(err)
			}

		default:
			fmt.Println("command does'nt exist...")
			choice = -1
		}

	}

}
