package db

/*
Seedable represents an structure which can seed db script
*/
type Seedable interface {
	DbName() string
	SeedScript() string
}
