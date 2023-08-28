package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	offset, index := 0, 1

	// start
	x, y := 0, 0
	for i := 0; i < n/2; i++ {
		for y < n-offset-1 {
			res[x][y] = index
			index++
			y++
		}

		for x < n-offset-1 {
			res[x][y] = index
			index++
			x++
		}

		for y > offset {
			res[x][y] = index
			index++
			y--
		}

		for x > offset {
			res[x][y] = index
			index++
			x--
		}

		x++
		y++
		offset++
	}

	// fill odd
	if n%2 == 1 {
		res[x][y] = index
	}

	return res
}

func main() {
	fmt.Println(generateMatrix(4))
}
