package main

func main() {
	Spiralize(5)
}

func ModifyArray(spiral [][]int) {
	n := len(spiral)
	IterDown(n, spiral)
	IterLeft(n, spiral)
	IterRight(n, spiral)
	IterUp(n, spiral)
	return
}
func IterRight(n int, spiral [][]int) {
	j := 0
	for i := 1; i <= n/2; i = i + 2 {
		for ; j < n-i; j++ {
			spiral[i][j] = 0
		}
		j = i
	}
}
func IterDown(n int, spiral [][]int) {
	initI := 1
	for j := n - 2; j >= n/2; j -= 2 {
		for i := initI; i < j+1; i++ {
			spiral[i][j] = 0
		}
		initI += 2
	}
}
func IterLeft(n int, spiral [][]int) {
	initJ := n - 2
	for i := n - 2; i >= n/2; i -= 2 {
		for j := initJ; j >= n-initJ-1; j-- {
			spiral[i][j] = 0
		}
		initJ -= 2
	}
}
func IterUp(n int, spiral [][]int) {
	initI := n - 2
	for j := 1; j <= n/2; j += 2 {
		for i := initI; i >= j+2; i-- {
			spiral[i][j] = 0
		}
		initI -= 2
	}
}

func Spiralize(size int) (res [][]int) {

	res = make([][]int, size)
	for i := range res {
		res[i] = make([]int, size)
	}
	for i := range res {
		for j := range res {
			res[i][j] = 1
		}
	}
	ModifyArray(res)
	return
}

type pathStep struct {
	times int
	dX    int
	dY    int
}

func SpiralizeACCEPTED(size int) [][]int {

	res := make([][]int, size)
	for i := range res {
		res[i] = make([]int, size)
	}
	res[0][0] = 1

	path := make([]pathStep, 0)
	path = append(path, pathStep{size - 1, 1, 0})

	for i := 0; path[i].times > 1; i++ {
		times := size - (i/2)*2 - 1
		if times > 0 {
			path = append(path, pathStep{times, -path[i].dY, path[i].dX})
		} else {
			break
		}
	}

	var x, y int
	for _, step := range path {
		for i := 0; i < step.times; i++ {
			x, y = x+step.dX, y+step.dY
			res[y][x] = 1
		}
	}

	return res
}
