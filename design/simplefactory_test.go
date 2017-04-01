package design

import (
	"testing"
)

func Test_operation(t *testing.T) {
	var a []Operation
	add := Divi{Data: []float64{1, 4, 5, 7, 9, 12}}
	a = append(a, add)
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
