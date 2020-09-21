package main

import "fmt"

func main() {

	//// 声明一个int8类型的数字,进行循环赋值并且输出.
	//// 没有赋初值.
	//var intarry [50]int8
	//for i := 0; i < 50; i++ {
	//
	//	intarry[i] = int8(i)
	//
	//}
	//
	//fmt.Println()
	//
	//// 进行循环遍历输出.
	//for i := 0; i < len(intarry); i++ {
	//
	//	// 取出数据进行遍历.
	//	fmt.Printf("取出第%v个数据: %v \n", i, intarry[i])
	//
	//}
	//
	//// 第二种循环体.
	//a := 0
	//for {
	//
	//	// 现在要通过当前循环输出所有数据.
	//	if a == 49 {
	//		fmt.Printf("取出第%v个数据: %v \n", a, intarry[a])
	//		break
	//	}
	//	a = a + 1
	//
	//}
	//
	//// 通过for rang键值循环.
	//str := "你好,GoLang"
	///**
	//for key,value := rang str
	//第一个参数是数组中的键值，第二个参数是值.
	//*/
	//for i, i2 := range str {
	//
	//	// 怎么输出其字符.
	//	// %c是输出UTF-8编码之后的的字符.
	//	fmt.Printf("key: %v  value: %c  \n", i, i2)
	//
	//}

	// 遇到输出结果不确定的时候测试输出%c%v这两种可能,但是也就这两种可能.
	//var array = [...]int{1,2,3,4,5,6,7}
	//for i := 0; i < len(array); i++ {
	//
	//	fmt.Printf("第%v次 %v \n",i,array[i])
	//
	//}

	//// 切片和数组定义很类似
	//var str = [] string{"1","2","#"}
	//// 切片保存的就是引用类型.
	//// 类似java对象引用,这个时候s2和str指向的是同一个内存地址.
	//s2:= str
	//s2[0] = "111"
	//fmt.Println(str)
	//// 切片对应Java中的集合.

	//var str = []string{"java","php","golang"}
	//
	//for k, v := range str {
	//
	//	fmt.Println(k,v)
	//	fmt.Printf("k %v v %v \n",k,v)
	//
	//}

	//从数组中取出某一些值那么取出来的就是切片.

	// 查看切面的容量和长度.
	//var slice = []int {0,1,2,3,4,5,6,7,8,9}
	//// 输出切面的容量和长度.
	//fmt.Printf("长度: %v 容量 %v",len(slice), cap(slice))
	//
	//// 从切片中定义切片.
	//// 其实就算这里取出两个值赋给a，但是a底层的数组的数据结构还是slice,只不过是将头指针移动到了下标为1的位置.
	//a := slice[1:3]
	//fmt.Printf("长度: %v 容量 %v",len(a), cap(a))
	//var a = add(12,8)
	//fmt.Println(a)
	go add(12,8)

}

/**
方法定义 : 方法参数是在括号里面,而方法的返回值是在参数后面.
这个和Java相反.
 */
func add(a int8, b int8) int8 {

	sum := a + b
	fmt.Println(sum)
	return sum

}
