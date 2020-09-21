package main

import (
	"fmt"
	"strconv"
)

/**
类型转换.
*/
func main() {

	// 使用fmt.Sprintf来进行格式转换.
	//a := 20
	//// 现在将20转换成字符串.
	//// 以%d格式输出的是字符串.
	//// 通过该方式可以将任意类型转换成字符串.
	//// 通过使用Sprigntf是将其他类型转换成String类型.
	//var f1 float64 = 25.661
	//str := fmt.Sprintf("%d", a)
	//fmt.Printf("%v - %T ",str,str)
	//
	//f := fmt.Sprintf("%f",f1)
	//fmt.Printf("\n %v - %T",f,f)

	// 现在看看怎么将字符串类型转换成其他类型.


	/*
		参数1：要转换的值
		参数2：格式化类型 'f'表示float，'b'表示二进制，‘e’表示 十进制
		参数3：表示保留的小数点，-1表示不对小数点格式化
		参数4：格式化的类型，传入64位 或者 32位
	*/

	var str string = "20"
	// 将字符串转换成int类型.
	// 第一个参数是要转换的值,第二个参数是要转换的进制,第三个参数是要转换的格式.
	var result, a = strconv.ParseInt(str,10,64)
	fmt.Printf("%v %T",result,result)
	fmt.Printf("%v %T",a,a)


}
