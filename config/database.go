package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Dbconnection() {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/bengkel")
	if err != nil {
		panic(err)
	}

	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)
	//db.SetConnMaxIdleTime(5 * time.Minute)
	//db.SetConnMaxLifetime(60 * time.Minute)

	DB = db

	log.Println("Database connected")

}
