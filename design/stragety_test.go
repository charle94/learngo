package design

import (
	"fmt"
	"testing"
)

func Test_stragety(t *testing.T) {
	context, err := GetCashContext("return", 10, 3)
	if err != nil {
		fmt.Println(err)
	}
	if context.GetResult(30) != 21 {
		t.Error("something wrong with test_stragety")
	}

}
