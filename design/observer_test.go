package design

import (
	"testing"
)

func Test_observer(t *testing.T) {
	s := ConcreteSubject{State: "watch"}
	o1 := ConcreteObserver{S: &s, State: "X", Name: "1"}
	o2 := ConcreteObserver{S: &s, State: "Y", Name: "2"}

	s.Attach(&o1)
	s.Attach(&o2)
	s.State = "watchover"
	s.Notify()
	if s.Ob[0].GetState() != "watchover" {
		t.Error("observer test wrong")
	}
}
