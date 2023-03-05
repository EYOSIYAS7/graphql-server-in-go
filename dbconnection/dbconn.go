package dbconnection

import (
	"fmt"
	"log"

	model "github.com/EYOSIYAS7/gptGraphql/Model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	dsn := "host=localhost port=5432 user=postgres dbname=test password=690177 sslmode=disable"
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=690177 sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("successfully connected to database")

	db.AutoMigrate(&model.Movie{})

	return db
}
