package main

import (
	"fmt"
	"math"
	"time"
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
	// a, _ := getData()
	// _, b := getData()
	// fmt.Println(a, b)

	var a byte = 'A'
	// 8进制写法: var a byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var a byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 'B'

	fmt.Println(a, b)

	str := [5]byte{104, 101, 108, 108, 111}

	fmt.Println(len(str))

	var country string = "hello, 中国"
	fmt.Println(len(country))

	arr()

	error(20)

	fmt.Println("everything is ok")

	fmt.Printf("now: %d", time.Now())

	switch i := 2; i * 4 {
	case 8:
		fmt.Println(i)
	default:
		fmt.Println("default")
	}

}

func error(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			// panic("oops")
		}
	}()

	var arr [10]int
	arr[x] = 88

}

func arr() {
	// var arr [3]int = [3]int{1, 2, 3}
	// fmt.Println(arr)
	// arr2 := [3]int{1, 2, 3}
	// fmt.Println(arr2)
	// arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(arr3)
	// fmt.Printf("%T", arr)
	// fmt.Printf("%T", arr2)
	// fmt.Printf("%T", arr3)

	// type arr4 [3]int

	// myarr := arr4{1, 2, 3}
	// fmt.Printf("%d 的类型是: %T", myarr, myarr)

	// arr5 := [4]int{2: 3, 1: 5}
	// fmt.Println(arr5)

	// arr := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// sliceArr := arr[1:3:4]
	// fmt.Println(sliceArr)

	myarr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("myarr 的长度为：%d, 容量为：%d\n", len(myarr), cap(myarr))

	mysli1 := myarr[1:3]
	fmt.Printf("mysli1 的长度为：%d, 容量为：%d, 类型为 %T\n", len(mysli1), cap(mysli1), mysli1)
	fmt.Println(mysli1)

	mysli2 := myarr[1:3:4]
	fmt.Printf("mysli2 的长度为：%d, 容量为：%d, 类型为 %T\n", len(mysli2), cap(mysli2), mysli2)
	fmt.Println(mysli2)

	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))

	arr := []int{1}
	arr = append(arr, 2)
	arr = append(arr, 3, 4)
	// ... 和 js 里面的 ... 一样，只不过写在后面
	arr = append(arr, []int{7, 8}...)
	arr = append([]int{0}, arr...)
	arr = append(arr[:5], append([]int{5, 6}, arr[5:]...)...)

	fmt.Println(arr)

	foo := []int{1, 2, 3}
	fmt.Println(len((foo)))
	fmt.Println(cap((foo)))

	foo = append(foo, 4, 5, 6, 7)
	fmt.Println(len((foo)))
	fmt.Println(cap((foo)))

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d, 其容量为: %d\n", myslice, len(myslice), cap(myslice))

	// 这里的 cap(4) 导致了截取了 numbers4 的 [5 6 7 8]
	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice为 %d, 其长度为: %d, 其容量为: %d\n", myslice, len(myslice), cap(myslice))

	// person := make(map[string]int)
	// person["name"] = 18
	// fmt.Println(person["name"])
	// fmt.Println(person["names"])

	fmt.Println(fibo(10))

	person := map[string]string{"name": "king", "age": "18"}

	if names, ok := person["names"]; ok {
		fmt.Println(names)
	} else {
		fmt.Println(ok)
	}

	for key, value := range person {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	name := "king"
	ptr := &name

	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(&name)
	fmt.Println(name)

	fmt.Println(ptr == &name)
	fmt.Println(*ptr == name)

	demoArr := [1]int{1}
	fmt.Println(demoArr)
	changeArrForPtr(&demoArr)
	fmt.Println(demoArr)

	demoArr2 := [2]int{1, 2}
	fmt.Println(demoArr2)
	changeArrForSls(demoArr2[:])
	fmt.Println(demoArr2)

	age := 18
	switch {
	case age > 6:
		fmt.Println(">6")
		fallthrough
	case age <= 6:
		fmt.Println("<=6")
	}

	hello := 1
	for hello <= 5 {
		fmt.Println(hello)
		hello++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("hello", i)
	}

	// br := 1

	// for {
	// 	if br >= 100 {
	// 		break
	// 		// return
	// 	}
	// 	fmt.Printf("hello %d", br)
	// 	br++
	// 	continue
	// }

	obj := map[string]int{"foo": 1, "bar": 2}

	for key, value := range obj {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	_arr := [...]int{1, 2, 3}

	for key, value := range _arr {
		fmt.Printf("key: %d, value: %d\n", key, value)
	}

	_sliceArr := []int{1, 2, 3, 4, 5}

	for _, value := range _sliceArr {
		if value == 3 {
			goto flag
		}

	flag:
		fmt.Println("stop")
	}

}

func judge(age int) bool {
	return false
}

func changeArrForPtr(arr *([1]int)) {
	(*arr)[0] = 100
}

func changeArrForSls(arr []int) {
	arr[0] = 100
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func int2bool(i int) bool {
	return i != 0
}

func fibo(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 || num == 2 {
		return 1
	} else {
		return fibo(num-1) + fibo(num-2)
	}
}
