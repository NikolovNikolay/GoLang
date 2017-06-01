package db

import (
	"database/sql"
)

/*
UnitOfWork ...
*/
type UnitOfWork struct {
	db             *sql.DB
	taskRepository *TaskRepository
}

/*
InitUnitOfWork ...
*/
func InitUnitOfWork(d *sql.DB) *UnitOfWork {
	if Uow == nil {
		Uow = &UnitOfWork{}
		Uow.db = d
	}
	return Uow
}

/*
ProvideTaskrepository ...
*/
func (uw *UnitOfWork) ProvideTaskrepository() *TaskRepository {
	if uw.taskRepository == nil {
		uw.taskRepository = &TaskRepository{db: uw.db}
	}

	return uw.taskRepository
}
