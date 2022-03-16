package belajar_golang_database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	username := os.Getenv("USERN")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DB")
	dataSource := username + ":" + password + "@tcp(" + host + ":3306)/" + database + "?parseTime=true"
	fmt.Println(dataSource)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
