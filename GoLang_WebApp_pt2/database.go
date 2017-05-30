package main

import (
	"database/sql"
	"fmt"
	"log"

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
	db, err = sql.Open(driverName, fmt.Sprintf("%s:%s@/%s", userName, password, dbName))
	if err != nil {
		panic(err.Error())
	}
	// // sql.DB should be long lived "defer" closes it once this function ends
	// defer db.Close()

	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection to DB tested successfully")
}

// DbGetAllTodos ...
func DbGetAllTodos() Todos {
	t := Todos{}

	var id int
	var name string

	rows, err := db.Query("select * from todos")

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return t
}
