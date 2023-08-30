package main

import (
	"fmt"
	"os"
)

var count int = 89

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
	Mkdir("./二叉树/二叉树的递归遍历/前序遍历")
	fmt.Println(count)
}
