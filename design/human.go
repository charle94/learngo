package design

import (
	"fmt"
)

type HunmanInterface interface {
	Hello()
}
type Human struct {
	Say string
}

func (h Human) Hello() {

	fmt.Println(h.Say)

}

type Student struct {
	Human
}

func GetHuman() HunmanInterface {
	s := Student{}
	s.Human.Say = "student"
	return s

}
