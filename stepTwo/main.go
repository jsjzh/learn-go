package main

import (
	"encoding/json"
	"log"
	"reflect"
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

func (person person) Demo() {

}

func (person person) logProfile() {
	log.Printf("名字：%s\n", person.name)
	log.Printf("年龄：%d\n", person.age)
	log.Printf("性别：%d\n", person.gender)
	log.Printf("工作：%s\n", person.work)
	log.Printf("公司：%s\n", person.company.companyName)
	log.Printf("公司：%s\n", person.companyName)
}

func (person *person) addAge() {
	person.age += 1
}

func (person *person) sayHello() {
	log.Println("hi, I'm developer")
}

func (person *person) assertion(i interface{}) {
	if _, ok := i.(string); !ok {
		log.Println("no no no")
	}

	log.Printf("i: %d", i)
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

	self.logProfile()
	self.addAge()
	self.logProfile()

	self1 := person{}
	self2 := new(person)
	self3 := &person{}

	log.Printf("self1's type: %T\n", self1)
	log.Printf("self1's name: %s\n", self1.name)
	log.Printf("self2's type: %T\n", self2)
	log.Printf("self2's name: %s\n", (*self2).name)
	log.Printf("self3's type: %T\n", self3)
	log.Printf("self3's name: %s\n", (*self3).name)

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
	log.Printf("该订单总共需要支付 %d 元", allPrice)
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
	log.Printf("%s\n", data1)

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
	log.Printf("%s\n", data2)
}

func typeAssertion() {

	var i interface{} = 10
	t1, ok := i.(int)
	log.Printf("%d-%t\n", t1, ok)

	log.Println("=====分隔线1=====")

	t2, ok := i.(string)
	log.Printf("%s-%t\n", t2, ok)

	log.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	log.Println(t3, "-", ok)

	// 这里是比较奇怪的，为什么一开始声明 k 是 interface{} 时候，断言是 interface{} 会出错
	// 但是给了 k 值之后，断言 interface{} 就对了
	// 难道是因为 k 的初始值为 nil
	// 哦，应该是的，interface{} 就和 ts 的 any 有点像
	// 可以被附上任何值
	log.Println("=====分隔线3=====")
	k = "10"
	t4, ok := k.(interface{})
	log.Printf("%s-%t\n", t4, ok)

	t5, ok := k.(int)
	log.Printf("%d-%t\n", t5, ok)

	var inter interface{} = 1
	tinter, ok := inter.(interface{})
	log.Printf("tinter's type: %T, tinter: %+v, ok: %t\n", tinter, tinter, ok)
	tinter, ok = inter.(int)
	log.Printf("tinter's type: %T, tinter: %+v, ok: %t\n", tinter, tinter, ok)
	tinter, ok = inter.(string)
	log.Printf("tinter's type: %T, tinter: %+v, ok: %t\n", tinter, tinter, ok)

	var kinter interface{}
	tkinter, ok := kinter.(interface{})
	log.Printf("tkinter's type: %T, tkinter: %+v, ok: %t\n", tkinter, tkinter, ok)
	kinter = "1"
	tkinter, ok = kinter.(int)
	log.Printf("tkinter's type: %T, tkinter: %+v, ok: %t\n", tkinter, tkinter, ok)
	tkinter, ok = kinter.(string)
	log.Printf("tkinter's type: %T, tkinter: %+v, ok: %t\n", tkinter, tkinter, ok)

	var developer IDev = &person{}
	_findType(developer)
}

func _findType(i interface{}) {
	switch x := i.(type) {
	case int:
		log.Println(x, "is int")
	case string:
		log.Println(x, "is string")
	case nil:
		log.Println(x, "is int")
	case IDev:
		log.Println(x, "is developer")
	default:
		log.Println(x, "no type matched")
	}
}

func someParamsFunc(params ...interface{}) {
	log.Printf("type %T, value %+v\n", params, params)
	for _, param := range params {
		log.Printf("type %T, value %+v\n", param, param)
		log.Println(param)
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

	log.Printf("person: %+v\n", person)
	log.Printf("arr: %+v\n", arr)
	log.Printf("sliceArr: %+v\n", sliceArr)
	log.Printf("demo: %+v\n", demo)
	log.Printf("demo2: %+v\n", demo2)

	slice := []int{1, 2, 3, 4, 5}
	var iSlice interface{}
	iSlice = slice
	// log.Println(iSlice[:])
	tiSlice, _ := iSlice.([]int)
	log.Println(tiSlice[:])
}

func reflection() {
	name := "king"

	log.Printf("name's type: %T, value: %s\n", name, name)

	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)

	log.Printf("t reflect's type's type: %T, value: %+v\n", t, t)
	log.Printf("v reflect's value's type: %T, value: %+v\n", v, v)

	n := v.Interface()

	log.Printf("n's type: %T, value: %s\n", n, n)

	v1 := reflect.ValueOf(&name)
	log.Printf("v1 reflect's value: %+v\n", v1)
	log.Printf("v1's CanSet: %t\n", v1.CanSet())
	v2 := v1.Elem()
	log.Printf("v2's CanSet: %t\n", v2.CanSet())

	p := person{
		name: "king",
	}

	vp := reflect.ValueOf(p)
	tp := reflect.TypeOf(p)

	for i := 0; i < vp.NumField(); i++ {
		log.Printf("%d Field: %+v", i, vp.Field(i))
	}

	for i := 0; i < tp.NumMethod(); i++ {
		log.Printf("%d Method: %+v", i, tp.Method(i).Name)
	}

}

func main() {
	defer func() {
		// body()
		// shopping()
		// tag()
		// typeAssertion()
		// goInterface()
		// reflection()
	}()
}
