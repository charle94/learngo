package design

import (
	"testing"
)

func Test_chain(t *testing.T) {
	h1 := new(ConcreteHandler_1)
	h2 := new(ConcreteHandler_2)
	h1.SetSuccessor(h2)
	h1.HandleRequest(19)
}
