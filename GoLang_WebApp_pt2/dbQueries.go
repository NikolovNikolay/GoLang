package main

import (
	"fmt"
)

const (
	columnID        = "id"
	columnName      = "name"
	columnCompleted = "completed"
	columnDue       = "due"

	tableTodos = "todos"

	queryGetAllTodos           = "SELECT " + columnID + "," + columnName + "," + columnCompleted + "," + columnDue + " FROM " + tableTodos
	queryPrepareInsertNewTodo  = "INSERT " + tableTodos + " SET " + columnName + "=?"
	queryPrepareGetTodoByID    = queryGetAllTodos + " WHERE " + columnID + "="
	queryPrepareDeleteTodoByID = "DELETE FROM " + tableTodos + " WHERE " + columnID + "="
)

// PrepareTodoQueryByID ...
func PrepareTodoQueryByID(id int64) string {
	return queryPrepareGetTodoByID + fmt.Sprintf("%v", id)
}

// PrepareTodoDeleteByID ...
func PrepareTodoDeleteByID(id int64) string {
	return queryPrepareDeleteTodoByID + fmt.Sprintf("%v", id)
}
