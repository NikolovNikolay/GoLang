package db

/*
UnitOfWorkProvider interface provides available repositories
*/
type UnitOfWorkProvider interface {
	ProvideTaskrepository() TaskRepository
}
