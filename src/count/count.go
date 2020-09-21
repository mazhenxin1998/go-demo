package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var sy sync.WaitGroup

// 定义互斥锁: 类比Java中的synchronized
var mutex sync.Mutex

/**
10个协程，每个协程添加100次, 最后结果应该是1000.
*/
func add() {

	// 延迟自执行函数.
	defer func() {

		error := recover()
		if error != nil {

			fmt.Printf("出现错误了. ")

		}

	}()

	// 每个协程计数100次.
	// lock出错了？
	mutex.Lock()
	for i := 0; i < 1000; i++ {
		// 尽量保证锁的代码块尽可能是最少的.
		count++
	}

	//sy.Done()
	mutex.Lock()

}

/**
模拟并发环境下count计数出现error的情况.
*/
func main() {

	// 开启10个协程计数.
	for i := 0; i < 100; i++ {

		//sy.Add(1)
		go add()

	}

	//sy.Wait()
	time.Sleep(time.Second * 10)
	// 最后这个结果是有错的,
	fmt.Printf("主线程结束: count最后结果: %v ", count)

}
