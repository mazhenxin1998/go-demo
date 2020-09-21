package main

import "fmt"

/**
断言.
*/

type Person interface {
	Say()
}

/**
第一个结构体.
*/
type Student struct {
	Name string
}

func (p *Student) Say() {

	fmt.Printf("Student Say方法发生了...")

}

func moreReturn() (int, int, string) {

	return 1, 1, "你好,GoLang"

}

/**
第二个结构体
*/
type Teacher struct {
	Name string
}

func (p Teacher) Say() {

	fmt.Printf("Teacher Say方法发生了. ")

}

/**
这里断言是不能使用的.
*/
func main() {

	a1, a2, a3 := moreReturn()
	fmt.Printf("第一个参数是: %v 第二个参数是: %v 第三个参数: %v", a1, a2, a3)

}
