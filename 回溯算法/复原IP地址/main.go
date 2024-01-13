package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	res := []string{}

	var dfs func(int, string, []string)
	dfs = func(index int, s string, tmpRes []string) {
		// 递归退出条件
		if len(tmpRes) == 4 && index == len(s) {
			res = append(res, tmpRes[0]+"."+tmpRes[1]+"."+tmpRes[2]+"."+tmpRes[3])
			return
		}

		// 递归
		for i := index; i < len(s); i++ {
			// 检查是否合法
			// 1. 0开头的数字只能是0
			// 2. 不能超过3位
			// 3. 数字不能大于255
			if i != index && s[index] == '0' {
				return
			}
			if i-index+1 > 3 {
				return
			}
			if v, _ := strconv.Atoi(s[index : i+1]); v > 255 {
				return
			}

			// 递归
			tmpRes = append(tmpRes, s[index:i+1])
			dfs(i+1, s, tmpRes)
			tmpRes = tmpRes[:len(tmpRes)-1]
		}
	}
	dfs(0, s, []string{})

	return res
}

func main() {
	// fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("101023"))
}
