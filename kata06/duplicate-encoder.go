package main

import (
	"fmt"
	"strings"
)

// The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.

// Examples
// "din"      =>  "((("
// "recede"   =>  "()()()"
// "Success"  =>  ")())())"
// "(( @"     =>  "))(("
// Notes
// Assertion messages may be unclear about what they display in some languages. If you read "...It Should encode XXX", the "XXX" is the expected result, not the input!

func main() {

	fmt.Println(DuplicateEncode("(( @"))
	fmt.Println(DuplicateEncode("Sucess"))
}

func DuplicateEncode(word string) string {
	var res string
	word = strings.ToLower(word)
	chars := make(map[string]int)

	for _, char := range word {
		chars[string(char)]++
	}

	for _, char := range word {
		val, _ := chars[string(char)]
		if val == 1 {
			res += "("
		} else {
			res += ")"
		}

	}

	return res
}

// same as mine but shorter syntax due to naked return and one liners
func DuplicateEncodeACCEPTED(word string) (r string) {
	word = strings.ToLower(word)

	t := map[rune]int{}
	for _, c := range word {
		t[c]++
	}

	for _, c := range word {
		if t[c] == 1 {
			r += "("
		} else {
			r += ")"
		}
	}

	return
}
