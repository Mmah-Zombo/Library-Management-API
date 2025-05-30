package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func migrateTable() error {
	log.Println("Running migrations.....")
	err := Db.AutoMigrate(&Book{})
	if err != nil {
		return err
	}
	log.Println("Models migrated successfully")
	return nil
}

func Init() (*gorm.DB, error) {
	godotenv.Load()
	dbUsr := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsr, dbPassword, dbHost, dbPort, dbName)

	var err error
	log.Println("Connecting to database.....")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to database")
		return nil, err
	}

	log.Println("Database connection successful")
	err = migrateTable()

	if err != nil {
		return nil, err
	}

	return Db, nil
}
