/**
Given a list of names, you need to organize each name within a slice based on its length.

Expected output:
[[] [] [Ava Mia] [Evan Neil Adam Matt Emma] [Emily Chloe]
[Martin Olivia Sophia Alexis] [Katrina Madison Abigail Addison Natalie]
[Isabella Samantha] [Elizabeth]]
*/

package main

import (
	"fmt"
)

var names = []string{
	"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

var nameMap = make(map[int][]string)

func main() {

	fillNameMap(names, nameMap)
	printMapVals(nameMap)
}

func fillNameMap(namesArr []string, mapParam map[int][]string) {

	for i := range names {

		n := names[i]
		l := len(n)

		if mapParam[l] == nil {
			newSlice := []string{n}
			mapParam[l] = newSlice
		} else {
			mapParam[l] = append(mapParam[l], n)
		}
	}
}

func printMapVals(mapParam map[int][]string) {
	for i := range mapParam {
		fmt.Printf("%q", mapParam[i])
	}
}
