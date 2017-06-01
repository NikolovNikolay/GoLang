package main

import (
	"exercise/GoLang_WebApp_pt2_ref/net/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
