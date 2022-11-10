package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("Connect to DB =>", db, "Error => ", err)
}
