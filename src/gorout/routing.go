package main

import (
	"fmt"
	"sync"
	"time"
)

/**
定义一个全局变量.
该变量用来控制协诚执行顺序.
需要注意的是,在Go中，如果主线执行完毕,那么协程也将会马上结束,不管是否执行完毕.
*/
var sy sync.WaitGroup

/**
协程通信读操作.
*/
func Write(ch chan int) {

	for i := 0; i < 10; i++ {

		fmt.Printf("写入操作: %v \n", i)
		// 将i写入管道中.
		ch <- i
		// 让当前协程休眠1秒.
		time.Sleep(time.Second)
	}

	// -1
	sy.Done()

}

/**
协程通信写操作.
*/
func read(ch chan int) {

	/**
	当读取比较快的时候会等待写入操作》
	*/
	for i := 0; i < 10; i++ {

		/**
		将ch中的值读入到result中.
		并且类型推导出result是一个int类型的值.
		*/
		result := <-ch
		fmt.Printf("读取到数据: %v \n", result)

	}

	sy.Done()

}

/**
管道是安全的，是一边写入，一边读取，当读取比较快的时候，会等待写入
*/
func main() {

	// 创建一个容量大小为10的管道进行数据传输.
	ch := make(chan int, 10)
	sy.Add(1)
	go Write(ch)
	sy.Add(1)
	go read(ch)
	sy.Wait()
	fmt.Printf("\n主线程执行完毕!")

}

