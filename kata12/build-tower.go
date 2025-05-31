package main

import (
	"fmt"
	"strings"
)

// Build Tower

// Build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors. A tower block is represented with "*" character.

// For example, a tower with 3 floors looks like this:

// [
//   "  *  ",
//   " *** ",
//   "*****"
// ]

// And a tower with 6 floors looks like this:

// [
//   "     *     ",
//   "    ***    ",
//   "   *****   ",
//   "  *******  ",
//   " ********* ",
//   "***********"
// ]

func main() {
	res := TowerBuilder(3)
	fmt.Println(res)
}

func AddChar(word string, n int, char string) string {
	for i := 0; i < n; i++ {
		word += char
	}
	return word
}

func TowerBuilder(nFloors int) []string {
	res := make([]string, nFloors)

	for i := 0; i < nFloors; i++ {
		// for i := range nFloors { -- not supported on codewars' golang version
		res[i] = AddChar(AddChar(AddChar(res[i], nFloors-(i+1), " "), ((nFloors*2)-1)-2*(nFloors-(i+1)), "*"), nFloors-(i+1), " ")
		fmt.Println(res[i])
	}

	return res
}

func TowerBuilderACCEPTED(nFloors int) []string {
	tow := make([]string, nFloors)
	for i := 0; i < nFloors; i++ {
		spaces := strings.Repeat(" ", nFloors-(i+1))
		blocks := strings.Repeat("*", i*2+1)
		tow[i] = spaces + blocks + spaces
	}
	return tow
}
