package main

import "fmt"

type company struct {
	companyName string
	companyAddr string
}

type Person struct {
	name, work string
	age        int
	gender     int8

	company

	father *Person
	mother *Person
}

func (person Person) FmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%d\n", person.gender)
	fmt.Printf("工作：%s\n", person.work)
	fmt.Printf("公司：%s\n", person.company.companyName)
	fmt.Printf("公司：%s\n", person.companyName)
}

func (person *Person) addAge() {
	person.age += 1
}

func main() {

	self := Person{
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

	self1 := Person{}
	self2 := new(Person)
	self3 := &Person{}

	fmt.Printf("self1's type: %T\n", self1)
	fmt.Printf("self1's name: %s\n", self1.name)
	fmt.Printf("self2's type: %T\n", self2)
	fmt.Printf("self2's name: %s\n", (*self2).name)
	fmt.Printf("self3's type: %T\n", self3)
	fmt.Printf("self3's name: %s\n", (*self3).name)

	obj := map[string]int{"age": 18}

	fmt.Printf("obj: %+v", obj)
}
