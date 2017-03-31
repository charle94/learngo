package learn

import (
	"fmt"
)

type List struct {
	Data     interface{}
	Nextnode *List
}

func myPrint(t interface{}) {
	fmt.Println(t)
}

//获取长度ok
func (l *List) ListLength() int {
	var i int = 0
	n := l
	for n.Nextnode != nil {
		i++
		n = n.Nextnode
	}
	return i + 1
}

//获取元素ok
func (l *List) GetEle(i int) (ele interface{}) {
	if i < 0 || i > l.ListLength() {
		return nil
	}
	cur := l
	j := 1
	for cur.Nextnode != nil {
		if j == i {
			return cur.Data
		}
		j++
		cur = cur.Nextnode

	}
	if cur != nil && j == i {
		return cur.Data
	}
	return nil
}

//插入变量到特定位置
func (l *List) ListInsert(i int, ele interface{}) bool {
	var s List //注意变量类型
	//可能存在无法插入到开头的情
	if i < 0 || i > l.ListLength() {
		return false
	}
	s.Data = ele
	p := l
	j := 1
	for j < i && p != nil {
		j++
		p = p.Nextnode
	}
	if p != nil && j <= i {
		s.Nextnode = p.Nextnode
		p.Nextnode = &s
	} else {
		return false
	}

	return true
}

//删除特定位置变量
func (l *List) ListDelete(i int) bool {
	p := l
	j := 1
	for j < i && p != nil {
		j++
		p = p.Nextnode
	}
	if p == nil || j > i {
		return false
	}
	p.Nextnode = p.Nextnode.Nextnode
	return true
}

//listall
func (l *List) ListAll() (all string) {
	p := l
	for p != nil {
		all += fmt.Sprintf("%v+", p.Data)
		p = p.Nextnode
	}
	return all
}
