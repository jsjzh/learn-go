package review

import (
	"log"
	"reflect"
)

// 这里主要是每天复习一下昨天重要的未完全理解的内容

func init() {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("error is: %v\n", err)
		}
	}()

	name := "king"
	ptr := &name

	log.Printf("ptr: %d\n", ptr)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Printf("arr: %+v\n", arr)
	log.Printf("arr's type: %T\n", arr)
	log.Printf("arr's len: %d\n", len(arr))
	log.Printf("arr's cap: %d\n", cap(arr))

	// start:end:len
	// [start, end) cap
	// cap = len(arr) - start
	sliceArr := arr[4:6:10]
	log.Printf("sliceArr: %+v\n", sliceArr)
	log.Printf("sliceArr's type: %T\n", sliceArr)
	log.Printf("sliceArr's len: %d\n", len(sliceArr))
	log.Printf("sliceArr's cap: %d\n", cap(sliceArr))

	// 切片是对数组的引用，如果用 append 等方法改变了 sliceArr，会对 arr 也产生影响
	// 但如果增加的长度超过了 arr，反而就不会对 arr 产生影响了，这个很神奇
	// 所以，没啥事儿就别用 arr 了，都用 sliceArr 吧
	appendSliceArr := append(sliceArr, 10, 11, 12, 13, 14)
	log.Printf("appendSliceArr: %+v\n", appendSliceArr)
	log.Printf("appendSliceArr's type: %T\n", appendSliceArr)
	log.Printf("appendSliceArr's len: %d\n", len(appendSliceArr))
	log.Printf("appendSliceArr's cap: %d\n", cap(appendSliceArr))
	log.Printf("arr: %+v\n", arr)
	log.Printf("arr's type: %T\n", arr)
	log.Printf("arr's len: %d\n", len(arr))
	log.Printf("arr's cap: %d\n", cap(arr))

	_appendSliceArr := append(sliceArr, 0)
	log.Printf("_appendSliceArr: %+v\n", _appendSliceArr)
	log.Printf("_appendSliceArr's type: %T\n", _appendSliceArr)
	log.Printf("_appendSliceArr's len: %d\n", len(_appendSliceArr))
	log.Printf("_appendSliceArr's cap: %d\n", cap(_appendSliceArr))
	log.Printf("arr: %+v\n", arr)
	log.Printf("arr's type: %T\n", arr)
	log.Printf("arr's len: %d\n", len(arr))
	log.Printf("arr's cap: %d\n", cap(arr))

	// map[string]int
	// 这个和 ts 有点像
	// type Demo = {[k:string]: number}
	// key 也是在 [] 里面
	person := map[string]int{"age": 18, "_age": 19}
	log.Printf("person: %+v\n", person)
	log.Printf("person's type: %T\n", person)

	// delete(person, "_age")

	// if names, ok := person["names"]; ok {
	// 	log.Printf("names: %d\n", names)
	// } else {
	// 	panic("error")
	// }

	count := 99
	log.Printf("count: %d\n", count)
	changeCount(&count)
	log.Printf("count: %d\n", count)

	changeArr := [...]int{1, 2, 3}
	log.Printf("changeArr: %+v\n", changeArr)
	changeArrForPtr(&changeArr)
	log.Printf("changeArr: %+v\n", changeArr)
	changeArr = [...]int{1, 2, 3}
	log.Printf("changeArr: %+v\n", changeArr)
	changeArrForSls(changeArr[:])
	log.Printf("changeArr: %+v\n", changeArr)

	if ok := judge(); ok {
		log.Println("yes")
	}

	switch num := getNumber(); num {
	case 1, 2, 3:
		log.Println("1, 2, 3")
	case num:
		log.Println("4, 5, 6")
	default:
		log.Println("10")
	}

	num := 6
	switch {
	case num > 5:
		log.Println(">5")
		fallthrough
	case num > 10:
		log.Println(">10")
	default:
		log.Println("hello")
	}

	foo := 1
	for foo < 10 {
		log.Printf("foo: %d\n", foo)
		foo++
	}

	while := 1
	for {
		if while > 10 {
			break
		}
		log.Printf("while: %d\n", while)
		while++
	}

	for key, value := range person {
		log.Printf("key: %s\n", key)
		log.Printf("value: %d\n", value)
	}

	me := &Person{
		name: "king",
		age:  18,

		Company: &Company{
			cname: "大搜车",
		},

		father: &Person{
			name: "king's father",
			age:  28,
		},
	}

	me.sayHello()
	me.addAge()
	me.sayHello()

	typeAssertion()
	reflex()
}

func reflex() {
	// name := map[any]any{
	// 	"name": "king",
	// }

	// t := reflect.TypeOf(name)
	// v := reflect.ValueOf(&name)

	// log.Println("t.Kind()", t.Kind())
	// log.Println("v.Kind()", v.Kind())
	// log.Println("v.Type()", v.Type())

	// log.Printf("t's type: %T, t's value: %+v", t, t)
	// log.Printf("v's type: %T, v's value: %+v", v, v)

	// log.Printf("CanSet: %t", v.CanSet())
	// v2 := v.Elem()
	// log.Printf("CanSet: %t", v2.CanSet())
	// log.Printf("name: %+v", name)
	// v2.SetString("hello king")
	// log.Printf("name: %+v", name)

	m := Profile{}

	t := reflect.TypeOf(&m)

	log.Println("&m Type: ", t)
	log.Println("&m Kind: ", t.Kind())

	log.Println("m Type: ", t.Elem())
	log.Println("m Kind: ", t.Elem().Kind())

	v := reflect.ValueOf(&m)

	log.Println("&m Type: ", v.Type())
	log.Println("&m Kind: ", v.Kind())

	log.Println("m Type: ", v.Elem().Type())
	log.Println("m Kind: ", v.Elem().Kind())
}

type Profile struct {
	name   string
	age    int
	gender string
}

func typeAssertion() {
	var foo any = 1
	tfoo, ok := foo.(any)
	log.Println(tfoo)
	log.Println(ok)

	tfoo, ok = foo.(int)
	log.Printf("tfoo's type: %T, tfoo's value: %+v", tfoo, tfoo)
	log.Printf("ok is %t", ok)

	tfoo, ok = foo.(bool)
	log.Printf("tfoo's type: %T, tfoo's value: %+v", tfoo, tfoo)
	log.Printf("ok is %t", ok)

	var bar any

	tbar, ok := bar.(any)
	log.Printf("tbar's type: %T, tbar's value: %+v", tbar, tbar)
	log.Printf("ok is %t", ok)

	bar = "hello"
	tbar, ok = bar.(string)
	log.Printf("tbar's type: %T, tbar's value: %+v", tbar, tbar)
	log.Printf("ok is %t", ok)

	tbar, ok = bar.(int)
	log.Printf("tbar's type: %T, tbar's value: %+v", tbar, tbar)
	log.Printf("ok is %t", ok)

	findType("123")
	arr := []int{1, 2, 3, 4, 5}
	findType(arr)
	obj := map[string]int{"foo": 123}
	findType(obj)
	person := &Person{name: "king"}
	findType(person)

	demo := map[any]any{
		"name":  123,
		123:     "hello",
		"hello": []int{1, 2, 3, 4, 5},
	}

	for key, value := range demo {
		findType(key)
		log.Printf("key: %+v, value: %+v", key, value)
	}

	var ddd any = []int{1, 2, 3, 4, 5}

	td := ddd.([]int)
	log.Println(td[:])

}

func findType(i any) {
	switch x := i.(type) {
	default:
		log.Printf("param's type: %T, param's value: %+v", x, x)
	}
}

type People interface {
	sayHello()
}

type Company struct {
	cname, caddr string
}

type Person struct {
	name, work string
	age        int
	gender     int8

	*Company

	father *Person
	mother *Person
}

func (person *Person) sayHello() {
	log.Printf("name: %s\n", person.name)
	log.Printf("age: %d\n", person.age)
	log.Printf("cname: %s\n", person.cname)
	log.Printf("father's name: %s\n", person.father.name)
}

func (person *Person) addAge() {
	person.age++
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
