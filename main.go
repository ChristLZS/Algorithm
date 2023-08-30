package main

import (
	"fmt"
	"os"
)

var count int = 180

func Mkdir(dir string) {
	// 创建目录
	os.MkdirAll(dir, 0755)

	// 生成go.mod文件
	f, _ := os.Create(dir + "/go.mod")
	f.WriteString(fmt.Sprintf("module %d", count))
	count++
	f.Close()

	// 生成main.go文件
	f, _ = os.Create(dir + "/main.go")
	f.WriteString(fmt.Sprintf("package main\n\nfunc main() {}"))
	f.Close()

	// 打开 ./go.work 文件
	// 追加 dir 到 ./go.work 文件
	f, _ = os.OpenFile("./go.work", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(fmt.Sprintf("%s\n", dir))
	f.Close()
}

func main() {
	Mkdir("./回溯算法/组合")
	Mkdir("./回溯算法/组合优化")
	Mkdir("./回溯算法/组合总和III")
	Mkdir("./回溯算法/电话号码的字母组合")
	Mkdir("./回溯算法/组合总和")
	Mkdir("./回溯算法/组合总和II")
	Mkdir("./回溯算法/分割回文串")
	Mkdir("./回溯算法/复原IP地址")
	Mkdir("./回溯算法/子集")
	Mkdir("./回溯算法/子集II")
	Mkdir("./回溯算法/递增子序列")
	Mkdir("./回溯算法/全排列")
	Mkdir("./回溯算法/全排列II")
	Mkdir("./回溯算法/重新安排行程")
	Mkdir("./回溯算法/皇后")
	Mkdir("./回溯算法/解数独")

	Mkdir("./贪心算法/分发饼干")
	Mkdir("./贪心算法/摆动序列")
	Mkdir("./贪心算法/最大子序和")
	Mkdir("./贪心算法/买卖股票的最佳时机II")
	Mkdir("./贪心算法/跳跃游戏")
	Mkdir("./贪心算法/跳跃游戏II")
	Mkdir("./贪心算法/K次取反后最大化的数组和")
	Mkdir("./贪心算法/加油站")
	Mkdir("./贪心算法/分发糖果")
	Mkdir("./贪心算法/柠檬水找零")
	Mkdir("./贪心算法/根据身高重建队列")
	Mkdir("./贪心算法/根据身高重建队列（续集）")
	Mkdir("./贪心算法/用最少数量的箭引爆气球")
	Mkdir("./贪心算法/无重叠区间")
	Mkdir("./贪心算法/划分字母区间")
	Mkdir("./贪心算法/合并区间")
	Mkdir("./贪心算法/单调递增的数字")
	Mkdir("./贪心算法/监控二叉树")

	Mkdir("./动态规划/斐波那契数")
	Mkdir("./动态规划/爬楼梯")
	Mkdir("./动态规划/使用最小花费爬楼梯")
	Mkdir("./动态规划/不同路径")
	Mkdir("./动态规划/不同路径II")
	Mkdir("./动态规划/整数拆分")
	Mkdir("./动态规划/不同的二叉搜索树")
	Mkdir("./动态规划/分割等和子集")
	Mkdir("./动态规划/最后一块石头的重量II")
	Mkdir("./动态规划/目标和")
	Mkdir("./动态规划/一和零")
	Mkdir("./动态规划/零钱兑换II")
	Mkdir("./动态规划/组合总和Ⅳ")
	Mkdir("./动态规划/爬楼梯进阶版")
	Mkdir("./动态规划/零钱兑换")
	Mkdir("./动态规划/完全平方数")
	Mkdir("./动态规划/单词拆分")
	Mkdir("./动态规划/打家劫舍")
	Mkdir("./动态规划/打家劫舍II")
	Mkdir("./动态规划/打家劫舍III")
	Mkdir("./动态规划/买卖股票的最佳时机")
	Mkdir("./动态规划/买卖股票的最佳时机II")
	Mkdir("./动态规划/买卖股票的最佳时机III")
	Mkdir("./动态规划/买卖股票的最佳时机IV")
	Mkdir("./动态规划/最佳买卖股票时机含冷冻期")
	Mkdir("./动态规划/买卖股票的最佳时机含手续费")
	Mkdir("./动态规划/最长递增子序列")
	Mkdir("./动态规划/最长连续递增序列")
	Mkdir("./动态规划/最长重复子数组")
	Mkdir("./动态规划/最长公共子序列")
	Mkdir("./动态规划/不相交的线")
	Mkdir("./动态规划/最大子序和")
	Mkdir("./动态规划/判断子序列")
	Mkdir("./动态规划/不同的子序列")
	Mkdir("./动态规划/两个字符串的删除操作")
	Mkdir("./动态规划/编辑距离")
	Mkdir("./动态规划/回文子串")
	Mkdir("./动态规划/最长回文子序列")

	Mkdir("./单调栈/每日温度")
	Mkdir("./单调栈/下一个更大元素I")
	Mkdir("./单调栈/下一个更大元素II")
	Mkdir("./单调栈/接雨水")
	Mkdir("./单调栈/柱状图中最大的矩形")

	Mkdir("./图论/所有可能的路径")
	Mkdir("./图论/岛屿数量")
	Mkdir("./图论/岛屿的最大面积")
	Mkdir("./图论/飞地的数量")
	Mkdir("./图论/被围绕的区域")
	Mkdir("./图论/太平洋大西洋水流问题")
	Mkdir("./图论/最大人工岛")
	Mkdir("./图论/单词接龙")
	Mkdir("./图论/钥匙和房间")
	Mkdir("./图论/岛屿的周长")
	Mkdir("./图论/并查集理论基础")
	Mkdir("./图论/寻找图中是否存在路径")
	Mkdir("./图论/冗余连接")
	Mkdir("./图论/冗余连接II")

	fmt.Println(count)
}
