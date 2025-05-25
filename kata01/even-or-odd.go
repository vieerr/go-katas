// 8 kyu

package main

import "fmt"

// type Node struct {
// 	key  int
// 	link *Node
// }

func main() {

	// a := Node{
	// 	key:  1,
	// 	link: nil,
	// }

	// b := Node{
	// 	key:  2,
	// 	link: nil,
	// }

	// a.link = &b
	// fmt.Println(a.key)
	// fmt.Println(a.link.key)
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func EvenOrOddBitwise(number int) string {
	return []string{"Even", "Odd"}[number&1]
}

func SplitStrings(str string) []string {
	fmt.Println(len(str))
	return []string{"asd", "asd"}
}
