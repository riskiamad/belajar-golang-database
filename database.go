package belajar_golang_database

import (
	"database/sql"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/joho/godotenv"
)

const projectDirName = "belajar-golang-database"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConnection() *sql.DB {
	loadEnv()

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
