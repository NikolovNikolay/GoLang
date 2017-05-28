package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/test.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing to file")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	f.Close()
}
