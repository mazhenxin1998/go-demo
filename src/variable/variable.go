package main

import "fmt"

/**
	练习变量的声明以及赋值和初始化.
*/
func main() {

	// 默认类型.
	var username = "张三"
	fmt.Print("var变量使用: " + username)
	// 使用带类型声明变量.
	var user string = "李四"
	// %v 表示安装原样输出.
	fmt.Printf("带类型赋值: %v", user)
	fmt.Println("")
	// 类型推导方式进行定义,简介明了,以后就这么使用.
	// 简洁明了.
	student := "我是学生通过类型推导赋值"
	fmt.Printf("类型推导赋值: %v\n",student)
	// 一次声明多个变量.
	// 一次声明的多个变量的类型一样.
	var s1,s2,s3 string
	s1 = "s1"
	s2 = "s2"
	s3 = "s3"
	fmt.Printf("一次对多个变量进行赋值: s1= %v ,s2=%v s3= %v \n",s1,s2,s3)

	// 一次对多个变量进行赋值,并且类型可以不一样.
	var (
		s4 string
		s5 int16
	)
	// 初始化.
	s4 = "string s4"
	s5 = 20
	// %d的作用是什么?
	fmt.Printf("一次对多个变量进行声明不同类型: s4= %d s5= %v \n",s4,s5)

	// 常量定义:
	// 常量在定义的时候必须进行赋初值.
	const i = 100
	fmt.Printf("输出常量值: %v" ,i)
	// iota常量计数器.
	const a = iota
	const b = iota
	fmt.Printf("使用iota常量计数器运行的结果: %v", b)
	fmt.Printf("使用iota常量计数器运行的结果: %v", a)





}
