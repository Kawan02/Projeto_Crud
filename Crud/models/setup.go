package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	database, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalln("Falha ao conectar no banco de dados.", err.Error())
	}

	database.AutoMigrate(&Book{})

	DB = database

}
