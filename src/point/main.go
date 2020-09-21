package main

import "fmt"

func main() {

	pointTest1()

}

/**
Go中的指针和C中的指针类似.
Go中在内存中开辟一个内存空间使用make关键字或者new关键字进行开辟.
*/
func pointTest1() {

	var a int = 8
	// 取变量指针的是用&进行取出变量的指针.
	// Go默认有类型推导.
	var p *int = &a
	// Go中有类型推导.
	var pp = &a
	fmt.Printf("%T", pp)
	// Go中默认类型是值类型,但是我要是用指针修改的话,那么就相当于修改原来的值.
	// 那么原来的值也会随之改变.
	*p = 20
	fmt.Printf("指针类型: %T", p)
	fmt.Println("\n", a)
	// 对指针取指: 如果变量a是一个指针, 那么变量a的内存地址上存储的就是一个内存地址.
	// 如果要访问指针内存所表示的内存地址的值.
	// 那么就在指针前面加上*表示读取指针表示的内存地址的值.
	// make和new的区别
	/**
	1. make只能用于切片，map,管道的初始化。
	2. new只能用于类型的内存分配.
	 */


}


