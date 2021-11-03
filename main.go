package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/KotaroYamazaki/go-clean-arch-example/cmd/app"
)

func main() {
	app.Run()
}
