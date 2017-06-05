package model

import "time"

/*
Task represents task object
*/
type Task struct {
	ID        int       `json:"id"`
	Subject   string    `json:"subject"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
