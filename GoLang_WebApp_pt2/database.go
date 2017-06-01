package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	userName   = "root"
	password   = "admin"
	dbName     = "go_todo_prj"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open(driverName, fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True", userName, password, dbName))
	checkError(err)

	// Test the connection to the database
	err = db.Ping()
	checkError(err)

	fmt.Println("Connection to DB tested successfully")
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
