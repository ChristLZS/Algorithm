package main

import "fmt"

func solveSudoku(board [][]byte) {
	var dfs func() bool
	dfs = func() bool {
		// 二维遍历
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				// 判断是否可以填充
				if board[i][j] != '.' {
					continue
				}

				for k := 1; k <= 9; k++ {
					board[i][j] = byte(k + '0')
					if check(board, i, j) {
						if dfs() {
							return true
						}
					}
					board[i][j] = '.'
				}
				return false
			}
		}
		return true
	}
	dfs()
}

// 检查位置[row][col]是否填充合法
func check(board [][]byte, row int, col int) bool {
	// 判断行
	for i := 0; i < len(board); i++ {
		if board[row][i] == board[row][col] && i != col {
			return false
		}
	}

	// 判断列
	for i := 0; i < len(board); i++ {
		if board[i][col] == board[row][col] && i != row {
			return false
		}
	}

	// 判断九宫格
	for i := row / 3 * 3; i < row/3*3+3; i++ {
		for j := col / 3 * 3; j < col/3*3+3; j++ {
			if board[i][j] == board[row][col] && i != row && j != col {
				return false
			}
		}
	}
	return true
}

func main() {
	// [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
	data := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(data)
	// 可视化打印数字
	for _, v := range data {
		for _, v1 := range v {
			fmt.Printf("%c ", v1)
		}
		fmt.Println()
	}

}
