package db

import (
	"database/sql"
	"exercise/GoLang_WebApp_pt2_ref/help"
	"fmt"

	// MqSql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	userName   = "root"
	password   = "admin"
	dbName     = "go_tasks_prj"
)

/*
Uow - the exported unit of work instance, that
contains all of the repositories for db CRUDs
*/
var Uow *UnitOfWork

func init() {
	sqlDB, err := sql.Open(driverName, fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True", userName, password, dbName))
	helpers.CheckError(err)

	// Test the connection to the database
	err = sqlDB.Ping()
	helpers.CheckError(err)

	// If ok initialize unit of work
	uw := InitUnitOfWork(sqlDB)

	/*
		Creating the tables in DB if needed
	*/
	seeder := NewDbSeeder(uw)
	seeder.Seed()

	fmt.Println("Connection to DB tested successfully")
}
