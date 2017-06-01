package db

import (
	"fmt"
)

// ConcatQueryWithParams ...
func ConcatQueryWithParams(q string, id int64) string {
	return q + fmt.Sprintf("%v", id)
}
