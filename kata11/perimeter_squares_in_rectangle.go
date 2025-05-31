// 5 kyu

package main

import "fmt"

// The drawing shows 6 squares the sides of which have a length of 1, 1, 2, 3, 5, 8. It's easy to see that the sum of the perimeters of these squares is : 4 * (1 + 1 + 2 + 3 + 5 + 8) = 4 * 20 = 80

// Could you give the sum of the perimeters of all the squares in a rectangle when there are n + 1 squares disposed in the same manner as in the drawing:

func main() {
	fmt.Println(Perimeter(5))
}

func fibonacci(n int) int {

	prev := 0
	next := 1
	res := 0
	for i := 0; i < n+1; i++ {

		res += next
		tmp := next
		next += prev
		prev = tmp
	}
	return res

}

func Perimeter(n int) int {
	return 4 * fibonacci(n)
}

func PerimeterFAVORITE(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return 4 * (a + b - 1)
}
