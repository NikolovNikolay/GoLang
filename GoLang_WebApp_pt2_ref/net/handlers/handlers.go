package handlers

import (
	"encoding/json"
	"errors"
	"exercise/GoLang_WebApp_pt2_ref/help"
	"exercise/GoLang_WebApp_pt2_ref/model"
	"exercise/GoLang_WebApp_pt2_ref/net/response"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"exercise/GoLang_WebApp_pt2_ref/db"

	"github.com/gorilla/mux"
)

const (
	contentType        = "Content-Type"
	contentTypeAppJSON = "application/json;charset=UTF-8"
	taskIDKey          = "taskID"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TasksIndex - Gets all tasks from DB
func TasksIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(http.StatusOK)

	tasks, tasksError := db.Uow.ProvideTaskrepository().GetAll()
	if helpers.CheckError(tasksError) != nil {
		checkForServerError(w, tasksError)
		return
	}

	res := response.New(http.StatusOK, tasks, nil)
	checkForServerError(w, json.NewEncoder(w).Encode(res))
}

// TaskShow - Gets a specifiv task from DB by its ID and shows it
func TaskShow(w http.ResponseWriter, r *http.Request) {
	var todoID int64
	var res *response.Response
	todoID, err := extractTodoID(r)

	if helpers.CheckError(err) != nil {
		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(http.StatusBadRequest)

		res = response.New(
			http.StatusBadRequest,
			nil,
			err)
		checkForServerError(w, err)
		return
	}

	t, err := db.Uow.ProvideTaskrepository().Find(todoID)
	if helpers.CheckError(err) != nil {
		checkForServerError(w, err)
		return
	}

	var finalStatus int

	if t == nil {
		finalStatus = http.StatusNotFound
	} else if t.Subject != "" {
		finalStatus = http.StatusFound
	} else {
		finalStatus = http.StatusNotFound
	}

	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(finalStatus)

	res = &response.Response{
		Status:  finalStatus,
		Payload: t,
		Error:   ""}

	checkForServerError(w, json.NewEncoder(w).Encode(res))
}

// TaskCreate ...
func TaskCreate(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	var res *response.Response
	// First we read all of the sent body and then we continue.
	// we have to be cautious if someone sends us some big data (the limit here is 1 048 576 bytes)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if e := helpers.CheckError(err); e != nil {
		checkForServerError(w, e)
		return
	}

	if e := helpers.CheckError(r.Body.Close()); e != nil {
		checkForServerError(w, e)
		return
	}

	if err := json.Unmarshal(body, &task); err != nil {
		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(http.StatusUnprocessableEntity)

		checkForServerError(w, err)
		return
	}

	t, err := db.Uow.ProvideTaskrepository().Create(task)
	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(http.StatusCreated)

	res = &response.Response{
		Status:  http.StatusCreated,
		Payload: t,
		Error:   ""}

	checkForServerError(w, json.NewEncoder(w).Encode(res))
}

// TaskDelete ...
func TaskDelete(w http.ResponseWriter, r *http.Request) {
	var todoID int64
	var res *response.Response

	todoID, err := extractTodoID(r)
	if err != nil {
		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(http.StatusBadRequest)
		checkForServerError(w, err)
		return
	}

	afected, err := db.Uow.ProvideTaskrepository().Delete(todoID)
	if err != nil {
		checkForServerError(w, err)
		return
	}

	if !afected && err == nil {
		// we assume the record with passed ID does not exist in DB
		res = &response.Response{
			Status:  http.StatusNotFound,
			Payload: nil,
			Error:   ""}

		w.Header().Set(contentType, contentTypeAppJSON)
		w.WriteHeader(http.StatusOK)
		checkForServerError(w, json.NewEncoder(w).Encode(res))
		return
	}

	w.Header().Set(contentType, contentTypeAppJSON)
	w.WriteHeader(http.StatusOK)

	res = &response.Response{
		Status:  http.StatusOK,
		Payload: nil,
		Error:   ""}

	checkForServerError(w, json.NewEncoder(w).Encode(res))
}

func extractTodoID(r *http.Request) (int64, error) {
	var taskID int64
	vars := mux.Vars(r)
	taskIDStr := vars[taskIDKey]
	if taskIDStr == "" {
		return 0, errors.New("Object ID not provided")
	}

	if v, err := strconv.ParseInt(taskIDStr, 10, 64); err == nil {
		taskID = v
	} else {
		return 0, err
	}

	return taskID, nil
}

func checkForServerError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		r := response.New(
			http.StatusInternalServerError,
			nil,
			err)

		checkForServerError(w, json.NewEncoder(w).Encode(r))
	}
}
