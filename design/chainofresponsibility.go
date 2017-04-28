package design

import (
	"fmt"
)

//使多个对象都有机会处理请求
//从而避免请求的发送者与接收者之间的耦合关系
//将这个对象连成一条链，沿着该链传递请求，直到有一个对象处理他为止
//可以将choice换成具体的request结构体中存储大量信息
type Handler interface {
	HandleRequest(choice int)
	SetSuccessor(s Handler)
}
type ConcreteHandler_1 struct {
	s Handler
}

func (c *ConcreteHandler_1) SetSuccessor(s Handler) {
	c.s = s
}
func (c ConcreteHandler_1) HandleRequest(choice int) {
	if choice > 0 && choice < 10 {
		fmt.Println("concretehandler_1", choice)
	} else if c.s != nil {
		c.s.HandleRequest(choice)
	}
}

type ConcreteHandler_2 struct {
	s Handler
}

func (c *ConcreteHandler_2) SetSuccessor(s Handler) {
	c.s = s
}
func (c ConcreteHandler_2) HandleRequest(choice int) {
	if choice > 10 && choice < 20 {
		fmt.Println("concretehandler_2", choice)
	} else if c.s != nil {
		c.s.HandleRequest(choice)
	}
}
