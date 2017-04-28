package design

import (
	"testing"
)

func Test_me(t *testing.T) {
	var m = &ConcreteMediator{cmap: make(map[string]Colleague)}
	var c1 = ConcreteColleague{Name: "小明", m: m}
	var c2 = ConcreteColleague{Name: "小李", m: m}
	var c3 = ConcreteColleague{Name: "小花", m: m}
	var c4 = ConcreteColleague{Name: "小草", m: m}
	m.Add(c1)
	m.Add(c2)
	m.Add(c3)
	m.Add(c4)
	c1.Send("hello")
	c2.Send("小花是谁")
}
