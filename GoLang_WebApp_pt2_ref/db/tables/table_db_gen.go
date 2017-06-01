package tables

/*
DbTableGen represents a generic DB table
*/
type DbTableGen struct {
	Name   string
	Script string
}

/*
SeedScript provides DB table generate script
*/
func (dbt DbTableGen) SeedScript() string {
	return dbt.Script
}

func (dbt DbTableGen) DbName() string {
	return dbt.Name
}
