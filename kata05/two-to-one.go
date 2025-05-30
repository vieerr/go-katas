package main

import (
	"fmt"
	"slices"
)

func main() {
	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"
	fmt.Println(TwoToOne(a, b))
}

func TwoToOne(s1 string, s2 string) string {
	total := s1 + s2
	elements := make(map[string]string)
	for _, element := range total {
		char := string(element)
		elements[char] = char
	}

	var keys []string
	for k := range elements {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	fmt.Println(keys)
	fmt.Println(elements)
	return "hola"
}
