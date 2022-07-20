package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type company struct {
	companyName string
	companyAddr string
}

type person struct {
	name, work string
	age        int
	gender     int8

	company

	father *person
	mother *person
}

func (person person) FmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%d\n", person.gender)
	fmt.Printf("工作：%s\n", person.work)
	fmt.Printf("公司：%s\n", person.company.companyName)
	fmt.Printf("公司：%s\n", person.companyName)
}

func (person *person) addAge() {
	person.age += 1
}

func (person *person) sayHello() {
	fmt.Println("hi, I'm developer")
}

func (person *person) assertion(i interface{}) {
	if _, ok := i.(string); !ok {
		fmt.Println("no no no")
	}

	fmt.Printf("i: %d", i)
}

type IDev interface {
	sayHello()
}

func body() {
	self := person{
		name:   "king",
		age:    18,
		gender: 1,
		work:   "developer",
		company: company{
			companyName: "大搜车",
		},
	}

	self.FmtProfile()
	self.addAge()
	self.FmtProfile()

	self1 := person{}
	self2 := new(person)
	self3 := &person{}

	fmt.Printf("self1's type: %T\n", self1)
	fmt.Printf("self1's name: %s\n", self1.name)
	fmt.Printf("self2's type: %T\n", self2)
	fmt.Printf("self2's name: %s\n", (*self2).name)
	fmt.Printf("self3's type: %T\n", self3)
	fmt.Printf("self3's name: %s\n", (*self3).name)

	self.sayHello()
}

type IGood interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name     string
	quantity int
	price    int
}

func (phone *Phone) settleAccount() int {
	return phone.price * phone.quantity
}

func (phone *Phone) orderInfo() string {
	return strconv.Itoa(phone.quantity) + phone.name + strconv.Itoa(phone.settleAccount())
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

func (gift *FreeGift) settleAccount() int {
	return 0
}

func (gift *FreeGift) orderInfo() string {
	return strconv.Itoa(gift.quantity) + gift.name + strconv.Itoa(gift.settleAccount())
}

func calculateAllPrice(goods []IGood) int {
	var total int
	for _, good := range goods {
		total += good.settleAccount()
	}
	return total
}

func shopping() {
	iPhone := &(Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	})

	earphones := &(FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	})

	goods := []IGood{iPhone, iPhone, iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func tag() {

	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)

	if err != nil {
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	// ================

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)
}

func typeAssertion() {

	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	fmt.Println("=====分隔线1=====")

	t2, ok := i.(string)
	fmt.Printf("%s-%t\n", t2, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	// 这里是比较奇怪的，为什么一开始声明 k 是 interface{} 时候，断言是 interface{} 会出错
	// 但是给了 k 值之后，断言 interface{} 就对了
	// 难道是因为 k 的初始值为 nil
	// 哦，应该是的，interface{} 就和 ts 的 any 有点像
	// 可以被附上任何值
	fmt.Println("=====分隔线3=====")
	k = "10"
	t4, ok := k.(interface{})
	fmt.Printf("%s-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)

	var inter interface{} = 1
	tinter, ok := inter.(interface{})
	fmt.Printf("tinter's type: %T, tinter: %+v, ok: %t\n", tinter, tinter, ok)
	tinter, ok = inter.(int)
	fmt.Printf("tinter's type: %T, tinter: %+v, ok: %t\n", tinter, tinter, ok)
	tinter, ok = inter.(string)
	fmt.Printf("tinter's type: %T, tinter: %+v, ok: %t\n", tinter, tinter, ok)

	var kinter interface{}
	tkinter, ok := kinter.(interface{})
	fmt.Printf("tkinter's type: %T, tkinter: %+v, ok: %t\n", tkinter, tkinter, ok)
	kinter = "1"
	tkinter, ok = kinter.(int)
	fmt.Printf("tkinter's type: %T, tkinter: %+v, ok: %t\n", tkinter, tkinter, ok)
	tkinter, ok = kinter.(string)
	fmt.Printf("tkinter's type: %T, tkinter: %+v, ok: %t\n", tkinter, tkinter, ok)

	var developer IDev = &person{}
	_findType(developer)
}

func _findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is int")
	case IDev:
		fmt.Println(x, "is developer")
	default:
		fmt.Println(x, "no type matched")
	}
}

func someParamsFunc(params ...interface{}) {
	fmt.Printf("type %T, value %+v\n", params, params)
	for _, param := range params {
		fmt.Printf("type %T, value %+v\n", param, param)
		fmt.Println(param)
	}
}

func goInterface() {
	someParamsFunc(1, 2, 3)

	person := map[string]interface{}{
		"name": "king",
		"age":  18,
	}

	arr := [...]interface{}{1, 2, 3, "name", "king"}
	sliceArr := []interface{}{1, 2, 3, "name", "king"}

	type Demo struct {
		name interface{}
	}

	demo := &Demo{
		name: "king",
	}

	demo2 := &Demo{
		name: 110,
	}

	fmt.Printf("person: %+v\n", person)
	fmt.Printf("arr: %+v\n", arr)
	fmt.Printf("sliceArr: %+v\n", sliceArr)
	fmt.Printf("demo: %+v\n", demo)
	fmt.Printf("demo2: %+v\n", demo2)

	slice := []int{1, 2, 3, 4, 5}
	var iSlice interface{}
	iSlice = slice
	// fmt.Println(iSlice[:])
	tiSlice, _ := iSlice.([]int)
	fmt.Println(tiSlice[:])
}

func main() {
	defer func() {
		// body()
		// shopping()
		// tag()
		// typeAssertion()
		goInterface()

		a := 10
		b := interface{}(a)

		switch interface{}(b).(type) {
		case int:
			fmt.Println("参数的类型是 int")
		case string:
			fmt.Println("参数的类型是 string")
		}
	}()
}
