package learn

import (
	"fmt"
	"testing"
)

func Test_sort(t *testing.T) {
	data := &SeqSort{Data: []int{9, 1, 5, 7, 4, 8, 10, 12, 16, 11, 13}}
	data.QuickSort()
	fmt.Println(data.Data)
}
