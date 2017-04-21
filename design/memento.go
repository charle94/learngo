package design

//在不破坏封装性的前提下，捕获一个对象的内部状态
//病在该对象之外保存这个状态。以后就可以将对象恢复到原先保存的状态
type Origin interface {
	CreateMemento() Memento
	SetMemento(m Memento)
	Show() string
	Set(s string)
}
type Originator struct {
	State string
}

func (o *Originator) CreateMemento() Memento {
	return Memento{state: o.State}

}
func (o *Originator) SetMemento(m Memento) {
	o.State = m.state
}
func (o *Originator) Show() string {
	return o.State
}
func (o *Originator) Set(s string) {
	o.State = s
}

type Memento struct {
	state string
}
type Caretaker struct {
	m Memento
}
