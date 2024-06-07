package database

import (
	_ "modernc.org/sqlite"
	"database/sql"
	"fmt"
)

var db *sql.DB

func GetDb() *sql.DB {
	return db
}

func Init() {

	var err error
	db, err = sql.Open("sqlite", "local.db")
	if err != nil {
		fmt.Println(err)
		fmt.Println("===============")
		panic("No database connection")
	}
	db.SetMaxOpenConns(10)
	prepDatabaseTable()
}

func prepDatabaseTable() {
	createTableBooks := ` 
	 CREATE TABLE IF NOT EXISTS books(
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 title TEXT,
	 isbn TEXT,
	 author TEXT,
	 year TEXT
	)`

	_, err := db.Exec(createTableBooks)
	if err != nil {
		fmt.Println(err)
		fmt.Println("==================")
		panic("Cannot create table books")
	}
}
