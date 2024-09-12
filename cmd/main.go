package main

import (
	"fmt"
	email "gomail/Email"
	"gomail/config"
)

func main() {
	// conexão com .env
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = email.LoadEmail("emaildestinatário", config.EmailHost, config.EmailPort, config.UserEmail, config.PasswordEmail)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enviado!")
}
