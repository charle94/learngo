package design

import (
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
