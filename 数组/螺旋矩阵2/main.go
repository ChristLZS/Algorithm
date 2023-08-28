package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	} else if len(matrix) == 1 {
		return matrix[0]
	}

	res := make([]int, 0)
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}

		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}

		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}

		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{3}, {2}, {1}}))
}
