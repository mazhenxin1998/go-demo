package main

import (
	"fmt"
	"sync"
)

var count = 0
var sy sync.WaitGroup

/**
10个协程，每个协程添加100次, 最后结果应该是1000.
*/
func add() {

	// 每个协程计数100次.
	for i := 0; i < 1000; i++ {

		count++
	}

	sy.Done()
}

/**
模拟并发环境下count计数出现error的情况.
*/
func main() {

	// 开启10个协程计数.
	for i := 0; i < 100; i++ {

		sy.Add(1)
		go add()

	}

	sy.Wait()
	// 最后这个结果是有错的,
	fmt.Printf("主线程结束: count最后结果: %v ",count)

}
