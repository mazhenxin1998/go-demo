package main1

import "fmt"

/**
结构体的定义和初始化.
*/
type Person struct {
	name string
	age  int
}

type Student struct {

	m map[string]string

}

/**
给结构体定义一个方法.
其中(p *Person) 表示该方法执行之后的结果印象到了了那个谁,那现在就规定这个谁是这个接受者.
*/
func (p *Person) initPerson(name string, age int) {

	fmt.Println("给结构体进行赋值的方法发生了. ")
	p.name = name
	p.age = age
}

/**
Go入口.
*/
func main() {

	var s = new(Student)
	// 如果现在你要给该m进行赋值
	// 那么首先应该对其进行分配内存空间
	// 现在需要将分配好的一个map内存空间赋值给s中的那个map.
	s.m = make(map[string]string)
	s.m["name"] = "张三"
	fmt.Printf("%v" , s)


	// 第一种初始化结构体.
	// 对于结构体可以使用new为一个结构体分配出一个内存.
	// 这里使用new关键字分配内存大小为Person.
	//var p = new(Person)
	//fmt.Printf("p的类型是: %T",p)
	//// 由于intiPerson方法在定义的时候使用的就是Person指针类型进行定义,那么在赋完值之后p的值也随之而改变.
	//p.initPerson("张三",18)
	//// 输出p的信息.
	//fmt.Printf("数据: %#v ",p)

	// 第二种初始化结构体的方法.

	//var p Person
	//p.age = 20
	//p.name = "我是张三"
	//fmt.Printf("%v 年龄: %v",p.name,p.age)

	// 第三那种方法进行初始化.
	// 使用&对结构体进行取地址操作，相当于对该结构体类型进行了一次new实例化操作
	// 从内存空间中开辟出一个Person大小的空间.并将该空间地址的引用赋值给p.

	//var p = &Person{}
	//p.name = "请叫我李四"
	//p.age = 88
	//fmt.Printf("%T \n",p)
	//fmt.Printf("%v 年龄: %v", p.name, p)

	// 第四种.
	// 类似JS中JSON进行格式化.
	// 如果使用直接赋值的话,那么最后一个属性
	//var p = Person{
	//
	//	name : "张三",
	//	age : 20,
	//}
	//
	//// 直接输出.
	//fmt.Printf("%T", p)


}
