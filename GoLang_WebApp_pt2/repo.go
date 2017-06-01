package main

import (
	"log"
	"time"
)

var currentID int
var todos Todos

// RepoFindTodo ...
func RepoFindTodo(id int64) Todo {
	stmt, err := db.Prepare(PrepareTodoQueryByID(id))
	checkError(err)

	rows, err := stmt.Query()
	checkError(err)
	defer rows.Close()

	t := Todo{}
	for rows.Next() {
		err := rows.Scan(&t.ID, &t.Name, &t.Completed, &t.Due)
		checkError(err)
		break
	}
	return t
}

// RepoCreateTodo ...
func RepoCreateTodo(name string) Todo {
	stmt, err := db.Prepare(queryPrepareInsertNewTodo)
	checkError(err)
	res, err := stmt.Exec(name)
	checkError(err)
	insID, err := res.LastInsertId()
	checkError(err)
	t := RepoFindTodo(insID)

	return t
}

// RepoDestroyTodo ...
func RepoDestroyTodo(id int64) (bool, error) {
	stmt, err := db.Prepare(PrepareTodoDeleteByID(id))
	checkError(err)
	_, err = stmt.Exec()
	checkError(err)

	return true, nil
}

// RepoGetAllTodo ...
func RepoGetAllTodo() Todos {
	t := Todos{}

	var (
		id        int
		name      string
		completed int
		due       time.Time
	)

	rows, err := db.Query(queryGetAllTodos)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &completed, &due)
		checkError(err)

		// parsing the Todo object
		newT := Todo{
			Name:      name,
			ID:        id,
			Completed: completed > 1,
			Due:       due}

		// adding to result collection
		t = append(t, newT)
		log.Println(id, name, completed, due)
	}

	return t
}
