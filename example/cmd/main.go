package main

import (
	"fmt"

	djan "github.com/alingse/djan-go"
	"github.com/alingse/djan-go/example/app"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var admin = djan.NewAdmin(db)

	admin.Register(&app.UserModel{})

	fmt.Println(admin)
}
