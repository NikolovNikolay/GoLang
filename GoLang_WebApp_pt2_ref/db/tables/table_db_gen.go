package tables

/*
DbTableGen represents a generic DB table
*/
type DbTableGen struct {
	Name   string
	Script string
}

/*
PassScript provides DB table generate script
*/
func (dbt DbTableGen) PassScript() string {
	return dbt.Script
}

/*
DbName provides the name of the table
*/
func (dbt DbTableGen) DbName() string {
	return dbt.Name
}
