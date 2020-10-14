package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sandbox/go-mailer-sandbox/cipher"
	"sandbox/go-mailer-sandbox/mailer"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("KEY")
	encrypt := os.Getenv("EMAIL_PASSWORD")
	password := cipher.Decrypt(encrypt, key)
	username := "aaronchen@jubo.health"
	from := "aaronchen@jubo.health"
	to := "pjchender@gmail.com"

	mailer.Send(username, from, to, password)
}
