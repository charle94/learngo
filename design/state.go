package design

import (
	"fmt"
)

//state 当一个对象的内在状态改变时允许改变其行为
//这个对象看起来好像是改变了其类
//解决当控制一个对象状态转换的条件表达式过于复杂时候的情况
//。把状态的判断转移到表示不同状态的一系列类中，可以将复杂的判断简化
/*type Context struct {
	S State
}

func (c *Context) Request() {
	c.S.Handle(c)
}

type State interface {
	Handle(c *Context)
}
type StateA struct {
}

func (s StateA) Handle(c *Context) {
	fmt.Println("a->b")
	c.S = new(StateB)
	//fmt.Println(c)
}

type StateB struct {
}

func (s StateB) Handle(c *Context) {
	//fmt.Println(c)
	fmt.Println("b->a")
	c.S = new(StateA)
}*/
//example
type Work struct {
	Hour    int
	S       State
	Finshed bool
}

func (w *Work) Time(t int) {
	fmt.Println(t)
	w.Hour = t
	w.S.Write(w)
}

type State interface {
	Write(w *Work)
}
type ForeNoon struct {
}

func (f ForeNoon) Write(w *Work) {
	if w.Hour < 12 {
		fmt.Println("当前工作状态良好")
	} else {
		w.S = Afternoon{}
		w.S.Write(w) //要记得再次调用否则只达到改变状态的目的
	}
}

type Afternoon struct {
}

func (a Afternoon) Write(w *Work) {
	if w.Hour < 17 {
		fmt.Println("当前处在下午精力充沛阶段")
	} else {
		w.S = Evening{}
		w.S.Write(w) //要记得再次调用否则只达到改变状态的目的

	}
}

type Evening struct {
}

func (e Evening) Write(w *Work) {
	if w.Finshed {
		fmt.Println("work done")
	} else {
		fmt.Println("加班中")
	}
}
