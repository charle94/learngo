package design

import (
	"fmt"
	"testing"
)

/*func Test_prototype(t *testing.T) {
	p1 := PrototypeInstance{Id: 1, TestStruct: &hello{Str: "hello"}}
	p2 := p1.Clone(2)
	if p1.TestStruct != p2.TestStruct {
		t.Error("error with prototype")
	}

}
*/
func Test_resume(t *testing.T) {
	x := Resume{Name: "x"}
	y := x.Clone()
	y.Name = "y"
	x.Name = "z"
	fmt.Println(y, x)

}
