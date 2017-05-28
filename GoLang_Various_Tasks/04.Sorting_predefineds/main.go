package main

import (
	"fmt"
	"sort"
)

func main() {
	letts := []string{"b", "f", "a", "q"}
	sort.Strings(letts)
	fmt.Printf("Strings: %q , sorted: %v \n", letts, sort.StringsAreSorted(letts))

	ints := []int{3, 5, 6, 1, 6, 9, 0}
	sort.Ints(ints)
	fmt.Printf("Ints: %d , sorted: %v", ints, sort.IntsAreSorted(ints))
}
