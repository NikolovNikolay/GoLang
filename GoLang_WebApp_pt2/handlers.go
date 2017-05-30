package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io"
	"io/ioutil"

	"strconv"

	"github.com/gorilla/mux"
)

const (
	contentType        = "Content-Type"
	contentTypeAppJSON = "application/json;charset=UTF-8"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex ...
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(http.StatusOK)
	checkForServerError(w, json.NewEncoder(w).Encode(DbGetAllTodos()))
}

// TodoShow ...
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoID int
	if v, err := strconv.Atoi(vars["todoID"]); err == nil {
		todoID = v
	} else {
		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(http.StatusBadRequest)
		checkForServerError(w, err)
		return
	}

	t := RepoFindTodo(todoID)
	w.Header().Set(contentType, contentTypeAppJSON)

	if t.Name != "" {
		w.WriteHeader(http.StatusFound)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	checkForServerError(w, json.NewEncoder(w).Encode(t))
}

// TodoCreate ...
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	// First we read all of the sent body and then we continue.
	// we have to be cautious if someone sends us some big data (the limit here is 1 048 576 bytes)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	checkError(err)
	checkError(r.Body.Close())

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(422) // unprocessable entity
		checkForServerError(w, err)
		return
	}

	t := RepoCreateTodo(todo)
	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(http.StatusCreated)

	checkForServerError(w, json.NewEncoder(w).Encode(t))
}

func checkForServerError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
