package design

import (
	"fmt"
	"testing"
)

func Test_iterator(t *testing.T) {
	a := new(ConcreteIterator)
	sli := []string{"one", "two", "three", "four", "five"}
	a.agg.CreateIterator(sli)
	l := a.Last()
	fmt.Println(l, a.current)
	for !a.IsDone() {
		fmt.Println(a.CurrentItem(), a.current)
		a.Pre()
		//fmt.Println(a.IsDone(), a.current)
	}

}
