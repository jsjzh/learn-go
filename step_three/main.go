package step_three

import (
	"learn-go/utils"
	"log"
)

type helloWorld interface {
	Hello()
}

type Demo struct {
	name string
}

func (demo *Demo) Hello() string {
	return demo.name
}

type HeiHeiHei helloWorld

func init() {
	log.Println("-------------------- from step three --------------------")
	utils.SayHello()
	utils.SayHi()

	heihei := &Demo{
		name: "king",
	}
	log.Println("heihei", heihei)

}
