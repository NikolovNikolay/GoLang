package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

	// // sql.DB should be long lived "defer" closes it once this function ends
	// defer db.Close()

	// Test the connection to the database
	err = db.Ping()
	checkError(err)

	fmt.Println("Connection to DB tested successfully")
}

// DbGetAllTodos ...
func DbGetAllTodos() Todos {
	t := Todos{}

	var (
		id        int
		name      string
		completed int
		due       time.Time
	)

	rows, err := db.Query("select * from todos")
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &completed, &due)
		checkError(err)

		// parsing the Todo object
		newT := Todo{
			Name:      name,
			ID:        id,
			Completed: completed > 1,
			Due:       due}

		// adding to result collection
		t = append(t, newT)
		log.Println(id, name, completed, due)
	}

	return t
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
