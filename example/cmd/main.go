package main

import (
	"fmt"

	"github.com/alingse/djan"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var admin = djan.NewAdmin(db)
	fmt.Println(admin)
}
