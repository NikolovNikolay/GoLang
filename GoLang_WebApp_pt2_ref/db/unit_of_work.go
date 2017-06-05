package db

import (
	"database/sql"
)

/*
UnitOfWork provides all available repositories
*/
type UnitOfWork struct {
	db             *sql.DB
	taskRepository *TaskRepository
}

/*
InitUnitOfWork initialized unit of work
*/
func InitUnitOfWork(d *sql.DB) *UnitOfWork {
	if Uow == nil {
		Uow = &UnitOfWork{db: d}
	}
	return Uow
}

/*
ProvideTaskrepository is a getter for tasks repository
*/
func (uw *UnitOfWork) ProvideTaskrepository() *TaskRepository {
	if uw.taskRepository == nil {
		uw.taskRepository = &TaskRepository{db: uw.db}
	}

	return uw.taskRepository
}
