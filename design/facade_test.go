package design

import (
	"testing"
)

func Test_facade(t *testing.T) {
	f := Facade{}
	f.MethodA()
	f.MethodB()
}
