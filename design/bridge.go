package design

//聚合表示一种弱的拥有关系，a可以包含b但是b不是a的一部分
//合成则是一种较强的拥有关系，体现了严格的部分和整体的关系，部分和整体的生命周期一样
//尽量使用合成聚合而不是继承
//桥接模式，将抽象部分与实现部分分离，使他们都可以独立变化
//实现系统有可能多角度分类，每一种分类都有可能变化，把这种多角度分离出来
//使其独立变化，减少耦合
import (
	"fmt"
)

type PhoneBrand interface {
	SetSoft(s Soft)
	Run()
}
type Soft interface {
	Run()
}
type Nokia struct {
	s Soft
}

func (n *Nokia) SetSoft(s Soft) {
	n.s = s
}
func (n Nokia) Run() {
	fmt.Println("nokia")
	n.s.Run()
}

type Game struct {
}

func (g Game) Run() {
	fmt.Println("手机游戏")
}
