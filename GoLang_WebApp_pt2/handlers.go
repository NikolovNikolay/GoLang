package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io"
	"io/ioutil"

	"strconv"

	"errors"

	"github.com/gorilla/mux"
)

const (
	contentType        = "Content-Type"
	contentTypeAppJSON = "application/json;charset=UTF-8"
	todoIDKey          = "todoID"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex ...
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(http.StatusOK)

	response := Response{
		Status:  http.StatusOK,
		Payload: RepoGetAllTodo(),
		Error:   nil,
	}

	checkForServerError(w, json.NewEncoder(w).Encode(response))
}

// TodoShow - Gets a specifiv Todo from DB by its ID and shows it
func TodoShow(w http.ResponseWriter, r *http.Request) {
	var todoID int64
	var response *Response
	if v, err := extractTodoID(r); err == nil {
		todoID = v
	} else {
		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(http.StatusBadRequest)

		response := &Response{
			Status:  http.StatusBadRequest,
			Payload: nil,
			Error:   err}
		checkForServerError(w, err)
		return
	}

	t := RepoFindTodo(todoID)
	w.Header().Set(contentType, contentTypeAppJSON)

	var finalStatus int
	if t.Name != "" {
		finalStatus = http.StatusFound
	} else {
		finalStatus = http.StatusNotFound
	}
	w.WriteHeader(finalStatus)

	response = &Response{
		Status:  finalStatus,
		Payload: t,
		Error:   nil}

	checkForServerError(w, json.NewEncoder(w).Encode(response))
}

// TodoCreate ...
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	var response *Response
	// First we read all of the sent body and then we continue.
	// we have to be cautious if someone sends us some big data (the limit here is 1 048 576 bytes)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	checkError(err)
	checkError(r.Body.Close())

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(http.StatusUnprocessableEntity)

		response = &Response{
			Status:  http.StatusUnprocessableEntity,
			Payload: nil,
			Error:   err}

		checkForServerError(w, err)
		return
	}

	t := RepoCreateTodo(todo.Name)
	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(http.StatusCreated)

	response = &Response{
		Status:  http.StatusCreated,
		Payload: t,
		Error:   nil}

	checkForServerError(w, json.NewEncoder(w).Encode(t))
}

// TodoDelete ...
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	var todoID int64
	var response *Response

	if v, err := extractTodoID(r); err == nil {
		todoID = v
	} else {
		response = &Response{
			Status:  http.StatusBadRequest,
			Error:   err,
			Payload: nil}

		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(http.StatusBadRequest)
		checkForServerError(w, err)
		return
	}

	_, err := RepoDestroyTodo(todoID)
	if err != nil {
		response = &Response{
			Status:  http.StatusBadRequest,
			Error:   err,
			Payload: nil}
		checkForServerError(w, err)
		return
	}

	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(http.StatusOK)

	j, err := json.Marshal(map[string]string{"ok": "true", "status": string(http.StatusOK)})
	checkForServerError(w, json.NewEncoder(w).Encode(j))
}

func extractTodoID(r *http.Request) (int64, error) {
	var todoID int64
	vars := mux.Vars(r)
	todoIDStr := vars[todoIDKey]
	if todoIDStr == "" {
		return 0, errors.New("Object ID not provided")
	}

	if v, err := strconv.ParseInt(todoIDStr, 10, 64); err == nil {
		todoID = v
	} else {
		return 0, err
	}

	return todoID, nil
}

func checkForServerError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		r := NewResponse(
			http.StatusInternalServerError,
			nil,
			err)

		checkForServerError(w, json.NewEncoder(w).Encode(r))
	}
}
