package main

import (
	"log"

	"github.com/KotaroYamazaki/go-clean-arch-sample/cmd/app"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	app.Run()
}
