package main

import (
	"fmt"

	djan "github.com/alingse/djan-go"
	"github.com/alingse/djan-go/example/app"
)

func initModel() {
	db := app.DB
	err := db.AutoMigrate(&app.UserModel{})
	if err != nil {
		panic(err)
	}
}

func main() {
	initModel()

	var admin = djan.NewAdmin(app.DB)
	admin.Register(&app.UserModel{})

	fmt.Println(admin)
}
