package design

import (
	"fmt"
)

//享元模式
//运用共享技术有效的支持大量细粒度的对象
//-------------------------------------
//享元模式可以避免大量非常相似类的开销
//在程序的设计中有时候需要生成大量细粒度
//的类实例来表示数据。如果能发现这些实例除了几个
//参数外基本上都是相同的，有时候就能够大幅度减少
//需要实例化的类的数量。如果将那些参数移到类实例的外面
//在方法调用时将他们传递进来，就可以通过共享大幅度减少单个实例的数目
//应用如棋盘中的棋子主要差别就是外部位置的不同
//所以实例化一个棋子对象传入不同的坐标来进行游戏
type Custom int
type FlyWeight interface {
	Operation(c Custom)
}
type ConcreteFlyWeight struct {
}

func (c ConcreteFlyWeight) Operation(cu Custom) {
	fmt.Println("具体flyweight", cu)
}

type UnshareConcreteFlyWeight struct {
}

func (u UnshareConcreteFlyWeight) Operation(c Custom) {
	fmt.Println("不共享的flyweight", c)
}

//工厂模式
var flyfactory = make(map[string]func() FlyWeight)

func FlyWeightFactory(s string) FlyWeight {
	return flyfactory[s]()
}
func RegisterFlyFactory(s string, f func() FlyWeight) {
	flyfactory[s] = f
}
func init() {
	RegisterFlyFactory("X", func() FlyWeight {
		return ConcreteFlyWeight{}
	})
	RegisterFlyFactory("U", func() FlyWeight {
		return UnshareConcreteFlyWeight{}
	})
}
