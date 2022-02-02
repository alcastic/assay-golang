package main

import (
	"fmt"
	"strings"
)

/*
A closure, also lexical closure or function closure,
is a technique for implementing lexically scoped name binding
in a language with first-class functions.
*/
func countWithIncrement(inc int) func() int {
	var count int
	return func() int {
		count += inc
		return count
	}
}

func main() {
	var countWithIncrementOfOne = countWithIncrement(1)
	var countWithIncrementOfTwo = countWithIncrement(2)

	fmt.Printf("countWithIncrementOfOne: %d\n", countWithIncrementOfOne())
	fmt.Printf("countWithIncrementOfTwo: %d\n", countWithIncrementOfTwo())
	fmt.Println(strings.Repeat("#", 10))
	fmt.Printf("countWithIncrementOfOne: %d\n", countWithIncrementOfOne())
	fmt.Printf("countWithIncrementOfTwo: %d\n", countWithIncrementOfTwo())
	fmt.Println(strings.Repeat("#", 10))
	fmt.Printf("countWithIncrementOfOne: %d\n", countWithIncrementOfOne())
	fmt.Printf("countWithIncrementOfTwo: %d\n", countWithIncrementOfTwo())
	fmt.Println(strings.Repeat("#", 10))
}
