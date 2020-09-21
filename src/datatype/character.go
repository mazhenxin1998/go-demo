package main

import "fmt"

/**
练习golang中的字符串.
golang中的字符串和Java中的不一样.
go中的字符串以原生数据类型出现, 其内部默认使用的是UTF-8编码, 字符串为("")里面的内容.
*/
func main() {

	//	s1 := "你好"
	//	var s2 string = "张三: "
	//	fmt.Println("练习字符串" + s1 + s2)
	//	// 使用字符串模板来定义多行字符串. 字符串模板和ES6中的类似.
	//	//第二行往后需要紧靠左侧,否则将会有空格显示.
	//	var s3 = `第一行
	//	第二行
	//第三行`
	//	fmt.Printf("输出使用字符串模板进行的 %v", s3)
	//	fmt.Println("字符串常见的操作. ")
	//	// UTF-8下的字符串编码占用三个字节.
	//	var str string = "你好: GoLang"
	//	// 使用go蘖枝字符串函数求字符串的长度.
	//	fmt.Printf("使用len()函数测试字符串的字节长度: %v", len(str))
	//	// 使用Sprint进行字符串拼接,其参数是当前所有需要的进行拼接的字符串.
	//	var result = fmt.Sprint(str, "x")
	//	fmt.Println("拼接之后的字符串: " + result)
	//	var ss string = "你好-张三-啊"
	//	// 现在对ss进行分割.
	//	// 和Java类似,返回一个数组.
	//	// Go中定义一个变量和Java中正好是相反的
	//	// Java中是先定义类型在定义变量.
	//	// golang中是先定义变量在定义类型. 而且类型也是可以不进行定义的》
	//	var arryStr []string = strings.Split(ss, "-")
	//	arry := strings.Split(ss, "-")
	//	fmt.Println("测试: " + arry[0])
	//	// %T 是输出指定参数的实际上的类型.
	//	fmt.Printf("转换之后的返回类型: %T \n", arryStr)
	//	// 判断是否包含.
	//	if strings.Contains(ss, "你好") {
	//
	//		fmt.Println("strings.Contains(`第一个参数: 表示该字符串中是否包含第二个参数指定的字符串`)")
	//
	//	}
	//
	//	// 判断字符串的前缀和后缀.
	//	// Java中是以str.endWith等来判断.
	//	// 后缀就不进行测试了, 和判断前缀类似.
	//	f1 := strings.HasPrefix(ss, "你")
	//	if f1 {
	//
	//		fmt.Printf("字符串%v是以前缀:%v开始的. \n",ss,"你")
	//
	//	}
	//
	//	// 这是ASCI吗进行输出.
	//	// 第一个参数可以使用_匿名内部类是使用.
	//	for i, i2 := range ss {
	//
	//		// 对字符串进行循环.
	//		// key表示当前字符串也就是value其实的下标,以字节开始计算. 一个汉字占用三个字节.
	//		// %c就是以原样输出,以我们直观感受.
	//		fmt.Printf("Key: %v Value: %c \n",i,i2)
	//
	//	}
	//
	//	// 字符串join操作.
	//	j1 := "j"
	//	j2 := "x"
	//	// 声明数组并且赋值需要将类型放到等号右侧.
	//	// %v 如果是汉子那么就输出ASCII吗.
	//	// 如果要输出汉子那么就只能使用%v.
	//	var strArrays = [3]string{"js",j1,j2}
	//	for i, array := range strArrays {
	//
	//		// %v表示输出什么?
	//		fmt.Printf("Key %v Value %v",i,array)
	//
	//	}
	//	// join是将一个字符串数组中的所有字符串按照规定的拼接符进行拼接.
	//	// Join不会使用.
	//	//strings.Join("-",strArrays)
	//	// 判断在字符串中的位置. 该位置在当前字符串表示的从0开始到当前字符串的字节大小.
	//	index := strings.Index(ss,"-")
	//	fmt.Printf("\n -第一次出现的意思是: %v", index)
	//
	//	// byte和rune类型.
	//	// TODO: 明天早上写byte和rune类型.
	//	// 所有的String类型操作,都放在strings下.
	//
	//	// 修改字符串.

	//var str string = "你好,GoLang"
	//var s1 string = "golang"
	//// 通过for循环控制字符串输出,那么输出的ASCII吗.
	//for i := 0; i < len(s1); i++ {
	//
	//	// 依次输出
	//	fmt.Printf("%c \n",s1[i])
	//
	//}
	//
	//fmt.Println("-----")
	//for i := 0; i < len(str); i++ {
	//
	//	// 这里对中文名字进行调用,会出现问题.
	//	fmt.Printf("%c \n",str[i])
	//
	//}
	//
	//for _, i2 := range str {
	//
	//	// 输出.
	//	fmt.Printf("%c",i2)
	//
	//}

	//要修改字符串，需要先将其转换成[]rune 或 []byte类型，
	//完成后在转换成string，无论哪种转换都会重新分配内存，并复制字节数组

	str := "你好,GoLang"
	//转换成rune类型.
	r := []rune(str)
	// 修改只能使用单个字符,而不能使用字符串进行修改.
	r[0] = '我'
	fmt.Printf("%c",r)

}
