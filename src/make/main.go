package main

import "fmt"

func main() {

	// 调用切片函数.
	//slice()
	makeMap()

}

func slice() {

	// 使用make来为切片分配空间.
	var s = make([]int8, 4, 4)
	// 切片赋值需要使用append
	// append用来进行扩容,用来删除
	// go中不能使用下标对数组进行扩容,如果要扩容那么就必须使用append方法.
	s = append(s, 1, 2, 3, 5, 6, 7, 8, 9, 6)
	fmt.Printf("值: %v  %T", s, s)

}

func makeMap() {

	// 使用make测试map.
	// map[KeyType]ValueType
	// go中的Map和Java中的Map是很类似的.
	var m = make(map[string]string)
	// 对map进行赋值.
	m["name"] = "张三"
	// 输出整个map值.
	fmt.Println("输出整个map值 " , m)
	fmt.Println("输出map中的值: ",m["name"])

	// make支持创建map的时候声明式赋值.
	var student = map[string]string{
		"username" : "张三",
		"age" : "20",
		"phone" : "xxxx",
	}

	fmt.Println("在创建map的时候就赋值: ",student["username"])
	// 遍历map.

	for k, v := range student {

		fmt.Printf("key----> %v  value---->%v  \n",k,v)

	}

	// 我们在获取某个值的时候，可以获取两个返回值.
	v,ok := student["age"]
	if ok {

		fmt.Println("从map中获取值的时候将会返回两个值，第一参数是返回的value第二个参数返回的是是否有该值",v)
		fmt.Println("如果存在键,那么就删除. ")
		delete(m,"age")
		// 调用方法进行输出map.
		travers(m)
	}

}

/**
对map进行遍历 .
 */
func travers(m map[string]string)  {

	fmt.Println("travers(map[string],string)方法执行了")
	for k, v := range m {

		fmt.Printf("key---> %v value---> %v \n",k,v)

	}

	
}
