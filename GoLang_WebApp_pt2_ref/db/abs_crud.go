package db

/*
Repositorier provides methods for repository CRUD operations
*/
type Repositorier interface {
	Find(id int64) interface{}
	Create(obj interface{}) interface{}
	Delete(id int64) (bool, error)
	GetAll() ([]interface{}, error)
}
