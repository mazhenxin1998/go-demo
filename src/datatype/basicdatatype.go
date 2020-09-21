package main

/**
先暂时将当前包改为非main,以后需要的时候在改回来.
*/
import (
	"fmt"
	"unsafe"
)

/**
基本数据类型: 1、整形 2、 浮点型 3、 布尔型 4、 字符串
*/
func main() {

	// int类型为有符号数.
	// 整形类型有 int8 int16 int32 int 64
	// int8表示占用一个字节八位.
	// 由于int8是占用一个字节,并且最高位表示符号位,那么int8的取值范围是-128-127
	// int类型的值默认是int8.
	i := 20
	// %v表示原样输出按照golang语言.
	fmt.Printf("查看int8类型占用多少位: %v", unsafe.Sizeof(i))
	// 强制类型转换.
	// 强制类型转换需要明确一个原理: 可以向下强制类型转换,但是不能向上强制类型转换.
	// 高位转低位需要注意一个问题就是会精度丢失问题.
	j := 130
	// 测试强制类型转换.
	// 这里为什么会变成-126?
	// TODO: 下面会介绍专用工具进行类型转换.
	fmt.Printf("\n测试下强制类型转换精度丢失问题: %v \n", int32(j))

	// 数字字面量.
	// 二进制 ob
	e := 0b00101101
	// 表示8进制.
	b := 0o337
	fmt.Printf("二进制数据: %v \n", e)
	fmt.Printf("八进制数据: %v \n", b)

	// 进制转换.
	// 原样数据是99
	number := 99
	// %v是原样输出. %d是十进制数据输出 .
	fmt.Printf("原样数据是: %v \n", number)
	// %d十进制数据输出.
	fmt.Printf("十进制数据输出: %v \n", number)
	// %o 以八进制数据输出. OCT英语表示的是八进制.
	fmt.Printf("八进制数据输出: %o \n", number)
	// %b 以二进制数据输出.
	fmt.Printf("以二进制数据输出: %b \n", number)
	// %x 表示以十六进制的数据进行输出. 十六进制: HEX
	fmt.Printf("以十六进制数据输出: %x \n", number)
	fmt.Println("-------------------")

	// GO支持两种浮点类型float32  float64
	// 标准浮点数输出 %f
	floatnumber := 22.123456789
	// %f打印小数点默认只打印小数点后6位.
	fmt.Printf("标准浮点数进行输出: %f \n", floatnumber)
	// 精确控制输出float小数点的后几位表示. 表示方法: %.nf 其中n表示要保留的小数点.
	fmt.Printf("精确控制输出float小数点后三位: %.3f \n", floatnumber)

	// bool类型.
	var flag bool = false
	// bool类型的另外几种表示方法.
	var f = false
	f1 := false
	if flag {

		fmt.Printf("bool类型进行输出: %v  %v %v \n", flag, f, f1)

	} else {

		fmt.Printf("bool类型进行输出: %v  %v %v \n", flag, f, f1)

	}

}
