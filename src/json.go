package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
定义一个结构体
结构体转换成JSON所以定义的变量名字必须是大写的.
注意的是: 结构体变量名字首字母必须是大写表示公有,否则json包是访问不到的,那么也就是说json包是访问不了到.
*/
type Student struct {
	Name string
	Age  int
	Sex  string
}

type Class struct {
	Title    string
	Students []Student
}

/**
嵌套结构体
*/
type Person struct {
	Id int
	St Student
}

func (p *Person) initPerson(Id int, St Student) {

	// 对其进行赋值.
	p.Id = Id
	// 首选应该为St开辟一个空间.
	p.St = St

}

func (p *Student) initStudent(name string, age int, sex string) {

	p.Name = name
	p.Age = age
	p.Sex = sex

}

func request(w http.ResponseWriter, r *http.Request) {

	// 嵌套结构体转换JSON.
	c := Class{

		Title:    "18软件二班",
		Students: make([]Student, 0), // 切片大小是可大可小.

	}

	// 循环赋值.
	for i := 0; i < 5; i++ {

		// 每次循环都需要声明变量.
		s := Student{

			Name: "李四",
			Age:  i,
			Sex:  "男",
		}

		// 现在需要将s赋值给c中的students.
		// append方法是Go内置的函数,其就是动态扩容.
		c.Students = append(c.Students, s)

	}

	// 输出c中的内容.
	fmt.Printf("%v", c)
	// 现在将其转换成JSON.

	bs, _ := json.Marshal(c)
	// 如果要输出,那么就需要将其转换成字符串.
	a := string(bs)
	fmt.Printf("%v",a)
	w.Write([]byte(a))


}

func main() {

	//// 将JSON转换成结构体.
	//var str = `{"Name": "你好,GoLang","Age"  : 20,"Sex"  : "男"}`
	//// 第一个参数是要转换的JSON,第二个地址是要将转换的结构体的值存放在那个结构体中.
	//var s = new(Student)
	//json.Unmarshal([]byte(str), &s)
	//// 转换成功. // 为什么没有显示值呢?
	//fmt.Printf("%v", s.Name)
	// 先对其进行赋值运算.
	//s := new(Student)
	//s.initStudent("张三", 88, "男")
	//p := new(Person)
	//fmt.Println(p)

	
}

/**
打开服务端监听一个端口,并且返回一个JSON对象.
*/
func init() {

	http.HandleFunc("/json", request)
	http.ListenAndServe("127.0.0.1:8080", nil)

}
