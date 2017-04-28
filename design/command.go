package design

import (
	"fmt"
)

//命令模式优点
//较容易设计一个命令队列
//在需要的情况下可以容易写入日志
//允许接受请求一方决定是否否决请求
//可以较容易实现对请求的撤销和重做
//将请求与具体的执行者分开
type Command interface {
	Execute()
}
type ConcreteCommand struct {
	r Receiver
}

func (c ConcreteCommand) Execute() {
	c.r.Action()
}

type Receiver interface {
	Action()
}
type ConcreteReceiver struct {
}

func (r ConcreteReceiver) Action() {
	fmt.Println("执行请求")
}

type Bab struct {
}

func (b Bab) Action() {
	fmt.Println("烧烤")
}

type Invoker struct {
	c []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.c = append(i.c, c)
}
func (i *Invoker) ExecuteCommand() {
	for _, v := range i.c {
		v.Execute()
	}
}
