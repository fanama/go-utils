package mail

import (
	"crypto/tls"
	"fmt"
	"log"

	gomail "gopkg.in/mail.v2"
)

type Configuration struct {
	Address     string `json:"address"`
	Password    string `json:"password"`
	Destination string `json:"destination"`
}

func (conf *Configuration) Init() (err error) {
	*conf, err = ReadConfig("mail")
	return err
}

func (conf *Configuration) Run() (err error) {
	err = conf.Init()

	if err != nil {
		log.Fatal(err)
	}

	choice := -1

	for choice != 0 {
		fmt.Println(">>>EmailSender<<<")
		fmt.Println(">1: Send an email")
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
			var message string
			fmt.Print("destination: ")
			fmt.Scan(&conf.Destination)
			fmt.Println("message: ")
			fmt.Scan(message)
			conf.SendMail(message)
		default:
			fmt.Println("command does'nt exist...")
			choice = -1
		}

	}

	return err
}

func (conf *Configuration) SendMail(message string) {
	m := gomail.NewMessage()
	// fmt.Println("sending mail to ", conf, "....")
	addresse := conf.Address
	password := conf.Password
	destination := conf.Destination

	// Set E-Mail sender
	m.SetHeader("From", addresse)

	// Set E-Mail receivers
	m.SetHeader("To", destination)

	// Set E-Mail subject
	m.SetHeader("Subject", "Nouveau Coli")

	// Set E-Mail body. You can set plain text or html with text/html

	m.SetBody("text/html", message)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, addresse, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}

}
