package db

import "exercise/GoLang_WebApp_pt2_ref/db/tables"
import "time"
import "fmt"

/*
DatabaseSeeder evaluates tables scripts for a DB
*/
type DatabaseSeeder struct {
	UnitOfWork *UnitOfWork
}

/*
NewDbSeeder constructs new database seeder
*/
func NewDbSeeder(uw *UnitOfWork) DatabaseSeeder {
	seeder := DatabaseSeeder{}
	seeder.UnitOfWork = uw
	return seeder
}

/*
Tables to seed
*/
func (dbs DatabaseSeeder) Tables() []Seedable {
	return []Seedable{
		tables.TableTasks,
	}
}

/*
DbName returns the name of the active DB
*/
func (dbs DatabaseSeeder) DbName() string {
	return dbName
}

/*
Seed evaluates the seed scripts for tables
*/
func (dbs DatabaseSeeder) Seed() {
	fmt.Println("Start DB seeding")
	go func() {
		for _, tab := range dbs.Tables() {
			if v, ok := tab.(Seedable); ok {
				_, e := dbs.UnitOfWork.db.Query(v.PassScript())
				if e != nil {
					fmt.Println(e.Error())
				} else {
					fmt.Println("Table " + v.DbName() + " created in DB")
				}
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
}
