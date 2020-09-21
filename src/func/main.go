package main

import "fmt"

/**
自定义函数类型.
*/
type calc func(a int, b int) int

/**
自定义函数类型作为参数进行传递值。
*/
type param func(a, b int) int

/**
自定义数据类型.
*/
type currentInt int

func main() {

	//fmt.Println("main...")
	//f1(1, 2, 3, 5)
	//fmt.Println("测试多返回值函数. ")
	//num, size := f2(1, 1, 2)
	//fmt.Printf("参数和未: %v 参数个数为: %v", num, size)
	//// 匿名函数类似JS中的一样,直接return一个函数即可.
	//// 测试匿名函数.
	//// 返回值是一个int类型的值.
	////   这里返回的是一个函数.
	//var result = anonymity(2)
	//fmt.Printf("\n测试返回值:%v %T", result(), result())
	//// 声明一个自定义的函数类型.
	//var f1 calc
	//// 而f1其实这个时候就是表示一个函数,类似Java中的Methodlei.
	//f1 = add
	//fmt.Printf("测试: %v  %T \n",f1,f1)
	//var result = f1(1,1)
	//fmt.Println(result)

	// 这里这个变量f就是一个函数的引用.
	//f := delete
	//fmt.Printf("类型: %T \n", f)
	//var a currentInt = 10
	//var b int = 10
	//fmt.Printf("测试: %T\n", a)
	//fmt.Printf("测试: %T\n", b)

	var ad param
	// 这里ad其实就是一个函数引用.
	p := ad
	var result = method(1, 2, p)
	fmt.Printf("以函数作为参数进行行为的传递: %v", result)

}

/**
创建可变参数的函数.
猜测: 应该和数组或者切片类似.
通常可变参数要作为函数的最后一个参数来接受.
*/
func f1(param ...int8) {

	// f1方法接受可变参数.
	for i := 0; i < len(param); i++ {

		fmt.Printf("可变参数第%v个 值为%v", i, param[i])

	}

}

/**
方法多返回值，Go语言中函数支持多返回值，同时还支持返回值命名，函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回
*/
func f2(param ...int) (num int, size int) {

	// 函数有两个返回值
	// num表示可变参数所有的和.
	// size表示可变参数的长度是多少.
	for i := 0; i < len(param); i++ {

		num = num + param[i]

	}
	size = len(param)
	// anonymity 
	// 直接使用return即可,那么返回的就是当前方法最终执行结果后的返回值的值.
	return

}

/**
测试匿名函数.
不需要返回值.
*/
func anonymity(o int) func() int {

	// 为什么这里匿名函数没有起到作用.
	switch o {
	case 0:
		return func() int {
			fmt.Println(0)
			return 1
		}
	default:
		return func() int {
			fmt.Println("默认分支.")
			return 0
		}

	}

}

func add(a int, b int) int {

	return a + b

}

func delete(a int, b int) int {

	return a - b

}

/**
方法作为参数.
*/
func t(a int, b int, sum func(int, int) int) int {

	// 这里就直接将第二个参数的方法进行返回.
	return 1
}

/**
实现以方法作为参数进行行为的传递.
*/
func method(a, b int, m param) int {

	// 这里这个m其实可以使用匿名函数进行操作.
	m = func(a, b int) int {
		return a + b
	}

	return m(a,b)

}
