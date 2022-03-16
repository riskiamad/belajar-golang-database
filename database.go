package belajar_golang_database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	userName := os.Getenv("USERN")
	pass := os.Getenv("PASSWORD")
	database := os.Getenv("DB")
	dataSource := userName + ":" + pass + "@tcp(" + host + ":3306)/" + database + "?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
