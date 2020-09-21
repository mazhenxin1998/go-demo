package main

import "fmt"

/**
Go主启动方法.
*/
func main() {

	// 先定义个异常监听.
	// 异常一定要放在defer中使用.
	// 并且在recover()函数是在一个匿名内部函数中进行延迟调用的.并且这个匿名函数是自执行的.
	defer func() {

		error := recover()
		if error != nil {

			fmt.Println(error)
		}

	}()
	testPanic()

	//testDefer()

}

/**
defer使用:
	在方法返回之前将会执行被defer 修饰的语句.
	并且执行defer修饰的语句也是从后往前进行执行.
*/
func testDefer() int {

	defer fmt.Println("延迟处理: testDefer方法发生了.")
	fmt.Println("testDefer方法级马上要返回了. ")
	return 0

}

/**
测试异常.
*/
func testPanic() {

	panic("抛出异常: ")

}
