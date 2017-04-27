package design

import (
	"fmt"
	"testing"
)

func Test_operation(t *testing.T) {
	o := GetOperation("divi", float64(1), float64(32))
	//var data = []float64{1, 4, 5, 7, 9, 12}
	//o := GetOperation("divi", data)
	//fmt.Println(o.GetResult(), data)
	var a []Operation
	a = append(a, o)
	a = append(a, Min{Data: []float64{1, 3, 6, 9}})
	a = append(a, Divi{Data: []float64{1, 3, 6, 9}})
	for i := 0; i < len(a); i++ {
		result := a[i].GetResult()
		if i == 1 {
			if -17 != result {
				t.Error("error with getting result")
			}
		}
	}
}

func Test_f(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	RegisterFactory("/", NewDivi)
	/*a := NewOperation("+", 1, 2, 3, 4)
	m := NewOperation("-", 1, 2, 3, 4)
	x := NewOperation("*", 1, 2, 3, 4)*/
	c := NewOperation("/", 1, 2, 3, 4)
	if c.GetResult() != 0.041666666666666664 {
		t.Error("简单工厂化方法错误")
	}
}
