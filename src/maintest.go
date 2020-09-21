package main

import "fmt"

func main() {

	// 使用go run main.go的时候一定要在当前文件所在的路径运行该命令.
	fmt.Print("Hello Word GoLang")
	// 使用printLn输出没给输出后面都会换行输出.
	fmt.Println("A","B","C")
	fmt.Println("")
	// go语言中的变量必须使用.
	var result = "aaa";
	fmt.Println("使用变量"+result)
	// 格式化输出 

}
