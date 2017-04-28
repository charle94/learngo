package design

import (
	"fmt"
)

//用一个中介对象来封装一系列的对象交互
//中介者使个对象不需要显式地相互引用
//而且可以独立的改变他们之间的交互
//系统中出现了多对多的交互复杂的对象群时

type Mediator interface {
	Send(message string, c Colleague)
	Add(c Colleague)
}
type Colleague interface {
	GetName() string
	Notify(message string)
}
type ConcreteColleague struct {
	m    Mediator
	Name string
}

func (c ConcreteColleague) GetName() string {
	return c.Name
}
func (c ConcreteColleague) Notify(message string) {
	fmt.Println(c.Name, "收到的信息", message)
}
func (c ConcreteColleague) Send(message string) {
	c.m.Send(message, c)
}

type ConcreteMediator struct {
	cmap map[string]Colleague
}

func (cm *ConcreteMediator) Add(co Colleague) {
	cm.cmap[co.GetName()] = co
}
func (cm ConcreteMediator) Send(message string, c Colleague) {
	for _, v := range cm.cmap {
		if v != c {
			v.Notify(message)
		}
	}
}
