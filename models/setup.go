package models

import (
	
  "gorm.io/gorm"
	"gorm.io/driver/mysql"
	

	
)

var DB *gorm.DB

func ConnecDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/golang_restapi_gin"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}