package db

import (
	"database/sql"
	"errors"
	"exercise/GoLang_WebApp_pt2_ref/db/tables"
	"exercise/GoLang_WebApp_pt2_ref/help"
	"exercise/GoLang_WebApp_pt2_ref/model"
	"log"
	"time"
)

/*
TaskRepository processes task objects in DB
*/
type TaskRepository struct {
	db *sql.DB
}

/*
NewTaskRepo ...
*/
func NewTaskRepo(db *sql.DB) *TaskRepository {
	repo := &TaskRepository{
		db: db}

	return repo
}

/*
Find a task in DB
*/
func (tr TaskRepository) Find(id int64) (model.Task, error) {
	stmt, err := tr.db.Prepare(ConcatQueryWithParams(tables.TableTasks.Consts.PrepareGetByID, id))
	if helpers.CheckError(err) != nil {
		return model.Task{}, err
	}

	rows, err := stmt.Query()
	if helpers.CheckError(err) != nil {
		return model.Task{}, err
	}
	defer rows.Close()

	t := model.Task{}
	for rows.Next() {
		err := rows.Scan(&t.ID, &t.Name, &t.Completed, &t.Due)
		if helpers.CheckError(err) != nil {
			return model.Task{}, err
		}
		break
	}
	return t, nil
}

/*
Create new Task object
*/
func (tr TaskRepository) Create(obj interface{}) (model.Task, error) {
	if obj == nil {
		return model.Task{}, errors.New("nil passed to create")
	}

	if _, ok := obj.(model.Task); !ok {
		return model.Task{}, errors.New("Invalid argument type")
	}

	taskToCreate := obj.(model.Task)

	if taskToCreate.Name == "" {
		return model.Task{}, errors.New("Provide at least task name")
	}
	stmt, err := tr.db.Prepare(tables.TableTasks.Consts.PrepareInsert)
	if helpers.CheckError(err) != nil {
		return taskToCreate, err
	}
	res, err := stmt.Exec(taskToCreate.Name)
	if helpers.CheckError(err) != nil {
		return taskToCreate, err
	}
	insID, err := res.LastInsertId()
	if helpers.CheckError(err) != nil {
		return taskToCreate, err
	}
	t, e := tr.Find(insID)
	if helpers.CheckError(e) != nil {
		return taskToCreate, e
	}

	return t, nil
}

/*
Delete a task from DB
*/
func (tr TaskRepository) Delete(id int64) (bool, error) {
	stmt, err := tr.db.Prepare(ConcatQueryWithParams(tables.TableTasks.Consts.PrepareDeleteByID, id))
	if helpers.CheckError(err) != nil {
		return false, err
	}
	_, err = stmt.Exec()
	if helpers.CheckError(err) != nil {
		return false, err
	}

	return true, nil
}

/*
GetAll retreaves all tasks from DB
*/
func (tr TaskRepository) GetAll() (model.Tasks, error) {
	t := model.Tasks{}

	var (
		id        int
		name      string
		completed int
		due       time.Time
	)

	rows, err := tr.db.Query(tables.TableTasks.Consts.QueryGetAll)
	if e := helpers.CheckError(err); e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &completed, &due)
		if helpers.CheckError(err) != nil {
			return nil, err
		}

		// parsing the Todo object
		newT := model.Task{
			Name:      name,
			ID:        id,
			Completed: completed > 1,
			Due:       due}

		// adding to result collection
		t = append(t, newT)
		log.Println(id, name, completed, due)
	}

	return t, nil
}
