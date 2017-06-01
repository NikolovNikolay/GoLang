package db

/*
UnitOfWorkProvider ...
*/
type UnitOfWorkProvider interface {
	ProvideTaskrepository() Repositorier
}
