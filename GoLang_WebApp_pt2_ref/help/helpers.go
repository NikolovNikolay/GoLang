package helpers

import (
	"fmt"
)

/*
CheckError logs the error and returns it back
*/
func CheckError(e error) error {
	if e != nil {
		fmt.Println(e)
	}

	return e
}
