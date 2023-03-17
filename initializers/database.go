package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {

	var err error
	// dsn := os.Getenv("DB_URL")
	dsn := "host=mahmud.db.elephantsql.com user=whwjqmsx password=WVdLH4uUpjxg-NOzfvBwVU-mVaWzN50H dbname=whwjqmsx port=5432 sslmode=disable"
	// fmt.Println("dsn")
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// fmt.Println("dsn")

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

}
