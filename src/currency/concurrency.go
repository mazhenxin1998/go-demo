package main

import (
	"fmt"
	"time"
)

/**
模拟下调用协程.
 */

func test() {

	fmt.Printf("我是协程...")

}


func main() {

	for i := 0; i < 10; i++ {

		go test()

	}

	// 明天把IDEA样式还原成最初始状态.
	// 结束.
	// 应该是毫秒吧.
	// 这里睡眠了多久?
	time.Sleep(10000)
	fmt.Println("协诚执行完毕!")

}
