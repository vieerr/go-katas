// 6 kyu

package main

import "fmt"

func main() {
	fmt.Println(Solution("Hola!"))
}

func Solution(str string) []string {
	res := []string{}

	for i := 0; i < len(str); i += 2 {
		var element string
		if i+1 >= len(str) {
			element = string(str[i]) + string("_")
		} else {
			element = string(str[i]) + string(str[i+1])
		}
		res = append(res, element)
	}

	return res
}

func AcceptedOne(str string) []string {
	var res []string

	if len(str)%2 != 0 {
		str += "_"
	}

	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}
