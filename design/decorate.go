package design

import (
	"fmt"
)

/**
 * example for decorate pattern
 */
/*import (
	"fmt"
)

type Compent interface {
	Operation()
}
type ConcreteCompent struct {
}

func (c ConcreteCompent) Operation() {
	fmt.Println("我的操作")
}

type Decorator struct {
	C Compent
}

func (d Decorator) Operation() {
	d.C.Operation()
}
func (d *Decorator) SetCompent(com Compent) {
	d.C = com
}

type DecoratorA struct {
	Decorator
}

func (d DecoratorA) Operation() {
	d.C.Operation()
	fmt.Println("A的操作")
}

type DecoratorB struct {
	Decorator
}

func (d DecoratorB) Operation() {
	d.C.Operation()
	fmt.Println("我是B的操作")
}
func client() {
	cc := ConcreteCompent{}
	d1 := DecoratorA{}
	d2 := DecoratorB{}
	d1.SetCompent(cc)
	d2.SetCompent(d1)
	//d2.SetCompent(d1)
	d2.Operation()

}*/
type Person interface {
	Show()
}
type PersonInstance struct {
	Name string
}

func (p PersonInstance) Show() {
	fmt.Println(p.Name)
}

type DecorateCloth struct {
	ClothName string
	P         Person
}

func (d *DecorateCloth) Decorate(P Person) {
	d.P = P
}
func (d DecorateCloth) Show() {
	fmt.Println("我穿上", d.ClothName)
	d.P.Show()
}
func GetCloth() {
	peson := PersonInstance{Name: "小明"}
	shirt := DecorateCloth{ClothName: "shirt"}
	paints := DecorateCloth{ClothName: "paints"}
	shirt.Decorate(peson)
	paints.Decorate(shirt)
	paints.Show()
}
