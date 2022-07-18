package main

import (
	"fmt"
	"math"
)

func print() {
	// fmt.Printf("")
	// %b    表示为二进制
	// %c    该值对应的unicode码值
	// %d    表示为十进制
	// %o    表示为八进制
	// %q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	// %x    表示为十六进制，使用a-f
	// %X    表示为十六进制，使用A-F
	// %U    表示为Unicode格式：U+1234，等价于"U+%04X"
	// %E    用科学计数法表示
	// %f    用浮点数表示
}

func declareVariables() {
	var (
		name string = "king"
		age  int    = 18
		// int 和 uint
		// uint 没有符号，可表示 2^8，0~255，256 个数字（uint8）
		// int 有符号，可表示 -128~127，也是 256 个数组（int8）
		sex  int8  = 127
		usex uint8 = 255
		// 这里默认会是 float64，因为系统是 64 位系统
		dot           = 0.12
		dot32 float32 = 0.12
		isDev bool    = true
		late  string
	)

	late = "hello"

	fmt.Println(name, age, sex, usex, dot, dot32, isDev, late)

	// := 自动推导，不用 var 声明
	auto := "king"

	fmt.Println(auto)

	autoName, autoAge, autoSex, autoIsDev := "autoKing", 18, 1, true

	fmt.Println(autoName, autoAge, autoSex, autoIsDev)

	var memoryAge int = 28
	// & 用来访问变量内存地址
	var memoryPath = &memoryAge

	fmt.Println(memoryAge)
	fmt.Println(memoryPath)
	// * 用来读取内存地址对应的变量
	fmt.Println(*memoryPath)

	// new 用来声明 int(也就是 type) 的零值，返回变量地址
	var foo *int = new(int)
	fmt.Println("foo memory path", foo)
	fmt.Println("foo", *foo)

	fmt.Println("MaxInt8", math.MaxInt8)
	fmt.Println("MaxUint8", math.MaxUint8)
	fmt.Println("MaxFloat32", math.MaxFloat32)
	fmt.Println("MaxFloat64", math.MaxFloat64)

	var myFloat float32 = 10000018

	fmt.Println("myFloat", myFloat)
	fmt.Println("myFloat + 1", myFloat+1)

	var myFloat2 float32 = 100000187
	fmt.Println("myFloat2", myFloat2)
	fmt.Println("myFloat2 + 1", myFloat2+1)
}

func getData() (int, int) {
	return 100, 200
}

func main() {
	declareVariables()
	// 匿名变量：_
	// 不分配内存，不占用内存空间
	// 可以多次声明
	a, _ := getData()
	_, b := getData()
	fmt.Println(a, b)

}
