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

func main() {

	nameMap := make(map[int][]string)

	for i := range names {

		n := names[i]
		l := len(n)

		if nameMap[l] == nil {
			newSlice := []string{n}
			nameMap[l] = newSlice
		} else {
			nameMap[l] = append(nameMap[l], n)
		}
	}

	for i := range nameMap {
		fmt.Printf("%q", nameMap[i])
	}
}
