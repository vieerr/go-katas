// 6 kyu

package main

import (
	"fmt"
	"strings"
)

var MORSE_CODE = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

func main() {
	// Morse code dictionary

	result := DecodeMorse(".... . -.--   .--- ..- -.. .")
	fmt.Println(result)
}

// not that proud of it
func DecodeMorse(morseCode string) string {

	var res string
	prev := 0
	morseCode += "  "
	var word string

	for i, element := range morseCode {
		if string(element) == " " {
			if prev == i {
				if len(word) != 0 {
					word += " "
				}
				res += word
				word = ""
			}
			word += MORSE_CODE[morseCode[prev:i]]
			prev = i + 1
		}
	}
	return strings.Trim(res, " ")
}

func DecodeMorseACCEPTED(morseCode string) (decoded string) {
	for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		for _, char := range strings.Split(word, " ") {
			decoded += MORSE_CODE[char]
		}
		decoded += " "
	}
	return strings.TrimSpace(decoded)
}
