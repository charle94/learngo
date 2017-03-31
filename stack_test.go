package learn

import (
	"fmt"
	"testing"
)

func myImport() {
	if false {
		fmt.Println("ok")
	}
}
func Test_List(t *testing.T) {
	//list测试--------------------------------
	var h, b, c List
	c.Data = 5
	c.Nextnode = nil
	b.Data = 3
	b.Nextnode = &c
	h.Data = 1
	h.Nextnode = &b

	//fmt.Println(b.ListLength())
	//fmt.Println(h.ListLength())
	h.ListInsert(1, 14)
	h.ListInsert(2, 13)
	h.ListInsert(1, 15)
	//fmt.Println(h.ListLength())
	//fmt.Println(h.ListAll())
	h.ListDelete(2)
	if h.ListLength() == 5 {
		t.Log("第一个通过")
	} else {
		t.Error("没有通过")
	}
}
func Benchmark_listInsert(b *testing.B) {
	h := List{Data: 1, Nextnode: nil}
	b.StopTimer()
	b.StartTimer()
	for i := 1; i < 10000; i++ {
		h.ListInsert(i, i+3)
	}
}
func Test_Stack(t *testing.T) {
	s := getSeqStack()
	if s.Top != -1 {
		t.Error("初始化失败")
	}
}
func Test_StackPush(t *testing.T) {
	s := getSeqStack()
	for i := 0; i < 10; i++ {
		ok := s.Push(i)
		if !ok {
			t.Error("插入失败")
		}
	}
}
func Test_getListStack(t *testing.T) {
	ls := getListStack()
	if ls.Count != 0 {
		t.Error("初始化队列失败")
	}
}
func Test_ListStachPush(t *testing.T) {
	ls := getListStack()
	for i := 0; i < 100; i++ {
		ls.Push(i)
	}
	if ls.Count != 100 {
		t.Error("出现错误")
	}
	//fmt.Println(ls)
}
