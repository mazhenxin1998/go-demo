package main

import (
	"fmt"
	"time"
)

func main() {

	//getTime()
	//setTimeOne()
	setTime()
}

func getTime() {

	time := time.Now()
	y := time.Year()
	m := time.Month()
	d := time.Day()
	fmt.Println(y, m, d)

}

/**
定时器
*/
func setTimeOne() {

	// 定义一个一秒中的定时器
	var time = time.NewTicker(time.Second)
	n := 0
	for i := range time.C {

		fmt.Println(i)
		n++
		if n >= 5 {

			// 终止定时器.
			time.Stop()
			return
		}

	}

}

func setTime() {

	n := 0
	for {

		// 休息1秒
		time.Sleep(time.Second)
		fmt.Println(n)
		n++
		if n > 5 {

			return
		}

	}

}
