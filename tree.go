package learn

import (
	"fmt"
)

type Tree struct {
	Data       interface{}
	Leftchild  *Tree
	Rightchild *Tree
}

func getTree() (tr *Tree) {
	return &Tree{
		Leftchild:  nil,
		Rightchild: nil,
	}
}
func (tr *Tree) PreOrderTraverse() {
	if tr == nil {
		return
	}
	fmt.Println(tr.Data)
	tr.Leftchild.PreOrderTraverse()
	tr.Rightchild.PreOrderTraverse()

}
func (tr *Tree) InOrderTraverse() {
	if tr == nil {
		return
	}
	tr.Leftchild.InOrderTraverse()
	fmt.Println(tr.Data)
	tr.Rightchild.InOrderTraverse()
}
func (tr *Tree) AfterOrderTraverse() {
	if tr == nil {
		return
	}
	tr.Leftchild.AfterOrderTraverse()
	tr.Rightchild.AfterOrderTraverse()
	fmt.Println(tr.Data)

}
