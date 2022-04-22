package main

import (
	"fmt"

	djan "github.com/alingse/djan-go"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var admin = djan.NewAdmin(db)
	fmt.Println(admin)
}
