package email

import (
	"bytes"
	"html/template"
	"strconv"

	"gopkg.in/gomail.v2"
)

func LoadEmail(to, host, port, username, password string) error {
	portconvert, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	// criar conexão
	dialer := gomail.NewDialer(host, portconvert, username, password)

	// criar mensagem
	msg := gomail.NewMessage()
	msg.SetHeader("From", "TesteGolang@gmail.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Teste Email Em Golang")
	msg.SetBody("text/html", getBody())

	// para anexa arquivos
	//msg.Attach("Email/images.png")

	// para adiconar imagem no corpo do email
	msg.Embed("Email/images.png", gomail.Rename("images.png"), gomail.SetHeader(map[string][]string{
		"Content-ID": {"<my-image>"},
	}))

	if err := dialer.DialAndSend(msg); err != nil {
		panic(err)
	}
	return nil
}

func getBody() string {
	t := template.Must(template.ParseFiles("Email/main.html"))
	var buff bytes.Buffer
	t.Execute(&buff, nil)
	return buff.String()
}
