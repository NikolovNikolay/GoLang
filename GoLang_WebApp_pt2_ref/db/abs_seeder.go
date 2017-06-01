package db

/*
Seeder evaluates the scripts for tables
*/
type Seeder interface {
	Tables() []Seedable
	DbName() string
	Seed()
}
