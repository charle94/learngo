package design

//观察者模式定义了一种一对多的
//依赖关系，这个主题对象在状态发生变化
//通知所有的观察者对象，使他们自动更新自己
type Observer interface {
	Update()
	GetState() string
	GetName() string
}

type Subject interface {
	Attach(o Observer)
	Delete(o Observer)
	Notify()
}
type ConcreteSubject struct {
	State string
	Ob    []Observer
}

func (c *ConcreteSubject) Attach(o Observer) {
	c.Ob = append(c.Ob, o)
}
func (c ConcreteSubject) Delete(o Observer) {

}
func (c *ConcreteSubject) Notify() {
	for _, v := range c.Ob {
		v.Update()
	}
}

type ConcreteObserver struct {
	State string
	Name  string
	S     *ConcreteSubject
}

func (c *ConcreteObserver) Update() {
	c.State = c.S.State
}
func (c ConcreteObserver) GetState() string {
	return c.State
}
func (c ConcreteObserver) GetName() string {
	return c.Name
}
