package main

import "fmt"

// 这里主要是每天复习一下昨天重要的未完全理解的内容

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error is: %v\n", err)
		}
	}()

	name := "king"
	ptr := &name

	fmt.Printf("ptr: %d\n", ptr)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("arr: %+v\n", arr)
	fmt.Printf("arr's type: %T\n", arr)
	fmt.Printf("arr's len: %d\n", len(arr))
	fmt.Printf("arr's cap: %d\n", cap(arr))

	// start:end:len
	// [start, end) cap
	// cap = len(arr) - start
	sliceArr := arr[4:6:10]
	fmt.Printf("sliceArr: %+v\n", sliceArr)
	fmt.Printf("sliceArr's type: %T\n", sliceArr)
	fmt.Printf("sliceArr's len: %d\n", len(sliceArr))
	fmt.Printf("sliceArr's cap: %d\n", cap(sliceArr))

	// 切片是对数组的引用，如果用 append 等方法改变了 sliceArr，会对 arr 也产生影响
	// 但如果增加的长度超过了 arr，反而就不会对 arr 产生影响了，这个很神奇
	// 所以，没啥事儿就别用 arr 了，都用 sliceArr 吧
	appendSliceArr := append(sliceArr, 10, 11, 12, 13, 14)
	fmt.Printf("appendSliceArr: %+v\n", appendSliceArr)
	fmt.Printf("appendSliceArr's type: %T\n", appendSliceArr)
	fmt.Printf("appendSliceArr's len: %d\n", len(appendSliceArr))
	fmt.Printf("appendSliceArr's cap: %d\n", cap(appendSliceArr))
	fmt.Printf("arr: %+v\n", arr)
	fmt.Printf("arr's type: %T\n", arr)
	fmt.Printf("arr's len: %d\n", len(arr))
	fmt.Printf("arr's cap: %d\n", cap(arr))

	_appendSliceArr := append(sliceArr, 0)
	fmt.Printf("_appendSliceArr: %+v\n", _appendSliceArr)
	fmt.Printf("_appendSliceArr's type: %T\n", _appendSliceArr)
	fmt.Printf("_appendSliceArr's len: %d\n", len(_appendSliceArr))
	fmt.Printf("_appendSliceArr's cap: %d\n", cap(_appendSliceArr))
	fmt.Printf("arr: %+v\n", arr)
	fmt.Printf("arr's type: %T\n", arr)
	fmt.Printf("arr's len: %d\n", len(arr))
	fmt.Printf("arr's cap: %d\n", cap(arr))

	// map[string]int
	// 这个和 ts 有点像
	// type Demo = {[k:string]: number}
	// key 也是在 [] 里面
	person := map[string]int{"age": 18, "_age": 19}
	fmt.Printf("person: %+v\n", person)
	fmt.Printf("person's type: %T\n", person)

	// delete(person, "_age")

	// if names, ok := person["names"]; ok {
	// 	fmt.Printf("names: %d\n", names)
	// } else {
	// 	panic("error")
	// }

	count := 99
	fmt.Printf("count: %d\n", count)
	changeCount(&count)
	fmt.Printf("count: %d\n", count)

	changeArr := [...]int{1, 2, 3}
	fmt.Printf("changeArr: %+v\n", changeArr)
	changeArrForPtr(&changeArr)
	fmt.Printf("changeArr: %+v\n", changeArr)
	changeArr = [...]int{1, 2, 3}
	fmt.Printf("changeArr: %+v\n", changeArr)
	changeArrForSls(changeArr[:])
	fmt.Printf("changeArr: %+v\n", changeArr)

	if ok := judge(); ok {
		fmt.Println("yes")
	}

	switch num := getNumber(); num {
	case 1, 2, 3:
		fmt.Println("1, 2, 3")
	case num:
		fmt.Println("4, 5, 6")
	default:
		fmt.Println("10")
	}

	num := 6
	switch {
	case num > 5:
		fmt.Println(">5")
		fallthrough
	case num > 10:
		fmt.Println(">10")
	default:
		fmt.Println("hello")
	}

	foo := 1
	for foo < 10 {
		fmt.Printf("foo: %d\n", foo)
		foo++
	}

	while := 1
	for {
		if while > 10 {
			break
		}
		fmt.Printf("while: %d\n", while)
		while++
	}

	for key, value := range person {
		fmt.Printf("key: %s\n", key)
		fmt.Printf("value: %d\n", value)
	}
}

func getNumber() int {
	return 10
}

func judge() bool {
	return 1 == 1
}

func changeCount(count *int) {
	*count = 100
}

func changeArrForPtr(arr *([3]int)) {
	(*arr)[1] = 100
}

func changeArrForSls(arr []int) {
	arr[1] = 100
}
