package main

import (
	"fmt"
	"sync"
)

/**
定义一个全局变量.
该变量用来控制协诚执行顺序.
需要注意的是,在Go中，如果主线执行完毕,那么协程也将会马上结束,不管是否执行完毕.
*/
var sy sync.WaitGroup

/**
定义一个全局变量,来接受计算之后的值.
*/

/**
该方法由四个携程计算12000累加.
*/
func computers(k int) {

	// 如果是0那么计算0 - 3000
	// 如果是1那么计算3001 - 6000
	var a = 0
	for i := (k - 1) * 3000; i < k*3000; i++ {

		a = a + i

	}
	fmt.Printf("当前协程计算结果是: %v",a)
	sy.Done()

}

/**
计算12000之和.利用四个协程.
*/
func main() {

	for i := 1; i <= 4; i++ {

		sy.Add(1)
		go computers(i)

	}

	// 主线程需要等待其他协程执行完毕才能释放.
	sy.Wait()
	fmt.Printf("主线程执行完毕. ")

}
