package step2

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

	var i interface{}
	var count int8
	fmt.Printf("type: %T\n", i)
	fmt.Printf("type: %T\n", count)

	t1, ok := i.(interface{})
	fmt.Println(ok)
	fmt.Println(t1)

	// fmt.Println("=====分隔线=====")

	// t2 := i.(string)
	// fmt.Println(t2)
}

func main() {
	defer func() {
		// body()
		// shopping()
		// tag()
		typeAssertion()
	}()
}
