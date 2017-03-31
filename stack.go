package learn

import (
	"fmt"
)

const (
	MAXSIZE = 10
)

type SeqStack struct {
	Data [MAXSIZE]interface{}
	Top  int
}
type ListStackNode struct {
	Data     interface{}
	Nextnode *ListStackNode
}
type ListStack struct {
	Count int
	Top   *ListStackNode
}

func (s SeqStack) String() string {
	str := ""
	for k, v := range s.Data {
		str += fmt.Sprintf("%v=%v\n", k, v)
	}
	return str
}
func getSeqStack() (s *SeqStack) {
	s = &SeqStack{Top: -1}
	return s

}
func (s *SeqStack) Push(ele interface{}) (ok bool) {
	if s.Top == MAXSIZE-1 {
		return false
	}
	s.Top++
	s.Data[s.Top] = ele
	return true
}
func (s *SeqStack) Pop() (ele interface{}, ok bool) {
	if s.Top == -1 {
		return nil, false
	}

	ele = s.Data[s.Top]
	s.Top--
	return ele, false
}

//-----------------------------
func getListStack() (ls *ListStack) {
	ls = &ListStack{
		Count: 0,
	}
	return ls
}
func (ls *ListStack) String() string {
	str := ""
	count := ls.getLength()
	for i := 0; i < count; i++ {
		str += fmt.Sprintf("%v=%v\n", i, ls.Pop())
	}
	return str
}
func (ls *ListStack) Push(ele interface{}) {
	var s ListStackNode
	s.Data = ele
	s.Nextnode = ls.Top
	ls.Top = &s
	ls.Count++
}
func (ls *ListStack) Pop() (ele interface{}) {
	if ls.Count == 0 {
		return nil
	}
	ele = ls.Top.Data
	ls.Top = ls.Top.Nextnode
	ls.Count--
	return ele
}
func (ls *ListStack) getLength() int {
	return ls.Count
}
