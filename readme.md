
# Envio de Email com Golang
É um projeto de envio de email utilizando a biblioteca Viper para fazer o gerenciamento do arquivo .env de forma dinâmica e GoMail para o setup de disparado de email.

## Baixando o Projetoy



```bash
  git clone https://github.com/TalissonSouzaDev/Enviodeemail_golang
```
```bash
  go mod tidy
```
### Configuração do .env
```bash
EMAIL_HOST=
EMAIL_PORT=
EMAIL_USER=
EMAIL_PASSWORD=
```
### Rodando o projeto
```bash
  go run cmd/main.go
```