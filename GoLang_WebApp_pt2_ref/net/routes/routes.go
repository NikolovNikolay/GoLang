package routes

import (
	"exercise/GoLang_WebApp_pt2_ref/net/handlers"
	"net/http"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"TasksIndex",
		"GET",
		"/tasks",
		handlers.TasksIndex,
	},
	Route{
		"TaskShow",
		"GET",
		"/tasks/{taskID}",
		handlers.TaskShow,
	},
	Route{
		"TaskCreate",
		"POST",
		"/tasks",
		handlers.TaskCreate,
	},
	Route{
		"TaskDelete",
		"POST",
		"/taskDelete/{taskID}",
		handlers.TaskDelete,
	},
}
