package main

import (
	"fmt"
	"sort"
)

type strBytes []string

func (sb strBytes) Less(i, j int) bool {
	return len(sb[i]) < len(sb[j])
}

func (sb strBytes) Swap(i, j int) {
	sb[i], sb[j] = sb[j], sb[i]
}

func (sb strBytes) Len() int {
	return len(sb)
}

func main() {
	cars := strBytes{"Mini", "GMC", "Mercedes-Benz", "Citroen"}
	fmt.Printf("Cars before sort: %q \n", cars)

	sort.Sort(strBytes(cars))
	fmt.Printf("Cars after sort: %q", cars)
}
