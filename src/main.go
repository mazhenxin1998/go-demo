package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

type Teacher struct {
	Id   int
	Name string
}

func (p *Teacher) TeacherInit(id int, name string) {

	p.Id = id
	p.Name = name

}

func reflectTest() {

	teacher := new(Teacher)
	teacher.Name = "张三"
	teacher.Id = 1
	// reflect.TypeOf返回参数中的类型结构.
	t := reflect.TypeOf(teacher)
	fmt.Printf("t.Elem().kind():  %v\n", t.Elem().Kind())
	fmt.Printf("类型名称: %v  类型种类: %v \n", t.Name(), t.Kind())
	// 获取第一个字段.
	f, ok := t.FieldByName("Name")
	if ok {

		fmt.Printf("Name的属性是: %v", f)

	}
	// 通过反射获取到的value是该结构体的地址的值.
	// reflect.ValueOf() 返回的是reflect.Value类型，其中包含了原始值的值信息，reflect.Value与原始值之间可以互相转换
	value := reflect.ValueOf(teacher)
	fmt.Printf("值: %v  \n", value)
	fmt.Println(t, value)

}

func readFile(filePath string) {

	// 解决异常的延迟自执行函数.
	defer func() {

		error := recover()
		if error != nil {

			fmt.Println("有异常发生了 ")

		}

	}()

	//reflectTest()
	// 对文件的操作.
	// 输出当前文件中的内容.
	file, error := os.Open("./json.go")
	if error == nil {

		// 正常执行操作.
		// 读取文件内容的缓冲区.
		cache := make([]byte, 1024)
		var str []byte
		for {

			_, er := file.Read(cache)
			if er == io.EOF {

				// 这个时候读取的内容都在cache缓存字节切片里面了.
				fmt.Printf("\n 读取完毕了. ")
				return

			}

			// 输出文件中的内容.
			str = append(str, cache...)
			fmt.Printf(string(cache))

		}

	} else {

		print("打开文件出现错误. ")

	}

}

func readFileByBufio(filePath string) {

	file, e := os.Open(filePath)
	defer file.Close()
	if e != nil {

		fmt.Println("打开文件出现错误. ")

	}

	// 通过创建bufio来读取文件.
	reader := bufio.NewReader(file)
	var fileStr string
	// 通过for循环按行读取.
	for {

		str, error := reader.ReadString('\n')
		if error == io.EOF {

			// 读取完成的时候也会有内容的.
			fileStr = fileStr + str
			fmt.Println("文件读取完毕. ")
			return

		}

		if error != nil {

			fmt.Println("读取我呢间出现错误. ")

		}

		fileStr = fileStr + str + "\n"
		// 输出.
		fmt.Print(fileStr)

	}

}

/**
- os.O_WRONLY：只读
- os.O_CREATE：创建
- os.O_RDONLY：只读
- os.O_RDWR：读写
- os.O_TRUNC：清空
- os.O_APPEND：追加
perm：文件权限，一个八进制数，r（读）04，w（写）02，x（执行）01
*/
func writeFile(path string) {

	// 先打开文件.
	file, _ := os.OpenFile(path, os.O_CREATE | os.O_RDWR | os.O_APPEND, 777)
	defer file.Close()
	str := "\r\n你好,GoLang"
	file.WriteString(str)


}

/**
学习反射.
*/
func main() {

	//writeFile("D:\\test\\test.txt")
	readFileByBufio("D:\\test\\test.txt")

}
