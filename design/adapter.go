package design

import (
	"fmt"
)

//definition 将一个类的接口转换成客户希望的另外一个接口
//。适配器模式可以是原本由于接口不兼容而不能在一起工作的
//那些类可以在一起工作

//系统的数据和行为都正确，但接口不符时
//考虑使用适配器，目的使控制范围之外的一个原有对象
//与某个接口匹配，适配器模式主要应用于一些希望复用一些现存的类，
//但接口与复用环境要求不一致的情况
//两个类做的事情相同或者相似，但是具有不同的接口要使用它
//
/*type Target interface {
	Request()
}
type Adaptee struct {
	//Target
}

func (a Adaptee) SpecialRequest() {
	fmt.Println("special request")
}

type Adapter struct {
	Adaptee
}

func (a Adapter) Request() {
	a.SpecialRequest()
}
func client() {
	var t Target = new(Adapter)
	t.Request()
}*/
type Player interface {
	Attack()
	Defense()
}
type Forward struct {
	Name string
}

func (f Forward) Attack() {
	fmt.Println("Forward attack", f.Name)
}
func (f Forward) Defense() {
	fmt.Println("Forward defense", f.Name)
}

type Center struct {
	Name string
}

func (c Center) Attack() {
	fmt.Println("center attack", c.Name)
}
func (c Center) Defense() {
	fmt.Println("center defense", c.Name)
}

type ForeignCenter struct {
	Name string
}

func (f ForeignCenter) 进攻() {
	fmt.Println("中文 进攻")
}
func (f ForeignCenter) 防御() {
	fmt.Println("中文 防御")
}

type Translator struct {
	Name string
	ForeignCenter
}

func (t Translator) Attack() {
	t.进攻()
	fmt.Println(t.Name)
}
func (t Translator) Defense() {
	t.防御()
	fmt.Println(t.Name)
}
