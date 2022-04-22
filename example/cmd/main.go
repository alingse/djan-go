package main

import (
	"log"
	"net/http"

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
	prefix := "/admin"
	http.HandleFunc(prefix, admin.Handler(prefix))

	log.Printf("serve at http://localhost:2345%s", prefix)
	http.ListenAndServe(":2345", nil)
}
