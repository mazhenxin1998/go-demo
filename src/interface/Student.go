package main

import (
	"fmt"
)

type Person interface {

	// 灭有返回值.
	Say(content string)
}

type Student struct {
	Name string
	Age  int
}

/**
实现结构体自己定义的方法.
*/
func (p *Student) StudentInit(name string, age int) {

	p.Name = name
	p.Age = age

}

/**
Go中对接口的实现,就是让结构体实现其接口中定义的方法就行.
*/
func (p *Student) Say(content string) {

	fmt.Printf("%v: %v", p.Name, content)

}

/**
空接口实践
*/
type NullInterface interface {
}

/**
可变参数的,并且参数是一个空接口.
也就是说该方法的参数是可以接受任何类型的结构体作为参数进行输出.
*/
func testInterface(a ...NullInterface) {

	fmt.Printf("利用空接口进行输出: %v", a)
}

/**
Map的值实现空接口.
*/

func main() {

	//// 给一个Student对象分配内存空间.
	//s := new(Student)
	//s.StudentInit("张三", 20)
	//
	//// 让当前student对象实例进行实现该接口.
	//// 声明一个接口的变量. 把要实现该接口的结构体的变量赋值给该接口的变量进行结构体实现接口操作.
	//var p Person = s
	//p.Say("你好, GoLang")

	testInterface("你好,GoLang")
	// 以空接口作为value来进行任何值的保存.
	var m = make(map[string]NullInterface)
	m["name"] = "张三"
	m["flag"] = true
	m["number"] = 11211
	fmt.Printf("查看Map中的值: %v  \n",m)

	// slice切片实现空接口.
	// 初始化切片内存空间的时候,需要考虑容量和长度大小.
	var s = make([]NullInterface,4,4)
	s[0] = "你好,GoLang"
	s[1] = true
	s[2] = 11211
	fmt.Printf("查看切片中的值: %v", s)

}
