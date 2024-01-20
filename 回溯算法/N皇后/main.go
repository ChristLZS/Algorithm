package main

import "fmt"

func solveNQueens(n int) [][]string {
	data := make([][]bool, n)
	for i := 0; i < n; i++ {
		data[i] = make([]bool, n)
	}
	res := make([][]string, 0)
	var dfs func(QNum int, row int, col int)
	dfs = func(QNum int, row int, col int) {
		// 递归终止条件
		if QNum == n {
			// 生成结果
			tmpRes := make([]string, n)
			for i := 0; i < n; i++ {
				s := make([]byte, n)
				for j := 0; j < n; j++ {
					if data[i][j] {
						s[j] = 'Q'
					} else {
						s[j] = '.'
					}
				}
				tmpRes[i] = string(s)
			}
			res = append(res, tmpRes)
			return
		}

		// 遍历列
		for j := col; j < n; j++ {
			// 判断是否可以放置皇后
			if check(data, row, j) {

				// 放置皇后，递归
				data[row][j] = true
				dfs(QNum+1, row+1, 0)
				data[row][j] = false
			}
		}
	}
	dfs(0, 0, 0)
	return res
}

func check(data [][]bool, row int, col int) bool {
	// 判断行
	for i := 0; i < len(data); i++ {
		if data[row][i] {
			return false
		}
	}

	// 判断列
	for i := 0; i < len(data); i++ {
		if data[i][col] {
			return false
		}
	}

	// 判断左上
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if data[i][j] {
			return false
		}
	}

	// 判断右上
	for i, j := row, col; i >= 0 && j < len(data); i, j = i-1, j+1 {
		if data[i][j] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(solveNQueens(4))
}
