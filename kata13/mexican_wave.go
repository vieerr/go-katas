// 6 kyu

package main

import (
	"fmt"
	"strings"
)

// Task

// In this simple Kata your task is to create a function that turns a string into a Mexican Wave. You will be passed a string and you must return an array of strings where an uppercase letter is a person standing up.
// Rules

// 1.  The input string will always consist of lowercase letters and spaces, but may be empty, in which case you must return an empty array. 2.  If the character in the string is whitespace then pass over it as if it was an empty seat
// Examples

// "hello" => ["Hello", "hEllo", "heLlo", "helLo", "hellO"]
// " s p a c e s " => [ " S p a c e s ", " s P a c e s ", " s p A c e s ", " s p a C e s ", " s p a c E s ", " s p a c e S "]

// Good luck and enjoy!

func main() {
	res := wave(" x yz")
	fmt.Printf("%q\n", res)
}

func wave(words string) []string {
	res := make([]string, len(strings.ReplaceAll(words, " ", "")))
	fmt.Println(len(res))
	j := 0
	for i := 0; i < len(words); i++ {
		// for i := range len(words) {
		wordArr := strings.Split(words, "")
		if wordArr[i] != " " {
			wordArr[i] = strings.ToUpper(wordArr[i])
			res[j] = strings.Join(wordArr, "")
			j++

		}
	}

	return res
}

func waveFAVORITE(words string) (wave []string) {
	wave = []string{} // leaving the array uninitialised would be better practice
	for i, c := range words {
		if c == ' ' {
			continue
		}
		upperC := string(c - 'a' + 'A')
		wave = append(wave, words[:i]+upperC+words[i+1:])
	}
	return
}
