package design

import (
	"fmt"
)

//使用了建造者模式用户
//只要指定建造的类型就可以得到他们
//二具体的建造细节就不需要知道了
//definition 将一个复杂对象的构建与他的表示分离
//使得同样的构建可以创建不同的表示
type PersonBuilder interface {
	BuildHead()
	BuildBody()
	BuildArmLeft()
	BuildArmRight()
	BuildLegLeft()
	BuildLegRight()
}
type PersonBuilderInstance struct {
}

func (p PersonBuilderInstance) BuildHead() {
	fmt.Println("创建头部")
}
func (p PersonBuilderInstance) BuildBody() {
	fmt.Println("创建身体")
}
func (p PersonBuilderInstance) BuildArmLeft() {
	fmt.Println("创建左胳膊")
}
func (p PersonBuilderInstance) BuildArmRight() {
	fmt.Println("创建右胳膊")
}
func (p PersonBuilderInstance) BuildLegLeft() {
	fmt.Println("创建左腿")
}
func (p PersonBuilderInstance) BuildLegRight() {
	fmt.Println("创建右腿")
}

type PersonDirector struct {
	pb PersonBuilder
}

func (p *PersonDirector) Set(pb PersonBuilder) {
	p.pb = pb
}
func (p PersonDirector) Build() {
	p.pb.BuildHead()
	p.pb.BuildBody()
	p.pb.BuildArmLeft()
	p.pb.BuildArmRight()
	p.pb.BuildLegLeft()
	p.pb.BuildLegRight()
}

type FatPerson struct {
	PersonBuilderInstance
}

func (f FatPerson) BuildBody() {
	fmt.Println("我是胖子的身体")
}
