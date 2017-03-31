package learn

import (
	"testing"
)

func Test_getTree(t *testing.T) {
	tr := getTree()
	tr.Data = 1
	p := tr
	for i := 2; i < 10; i++ {
		p.Rightchild = getTree()
		p.Rightchild.Data = i
		p.Leftchild = getTree()
		p.Leftchild.Data = i + 3
		p = p.Leftchild
	}
	tr.PreOrderTraverse()
	//fmt.Println("\n next \n")
	tr.InOrderTraverse()
	//fmt.Println("\n next \n")
	tr.AfterOrderTraverse()
}
