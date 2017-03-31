package learn

import (
	"fmt"
)

type BSTTree struct {
	Tree
	Leftchild  *BSTTree
	Rightchild *BSTTree
	Data       int
}

func (tr *BSTTree) PreOrderTraverse() {
	if tr == nil {
		return
	}
	fmt.Println(tr.Data)
	tr.Leftchild.PreOrderTraverse()
	tr.Rightchild.PreOrderTraverse()

}
func BinarySearch(from []int, target int) int {
	var low, high, mid int
	low = 0
	high = len(from) - 1
	for low <= high { //如果没有=会出现low=high的情况不返回值
		//fmt.Printf("riaoshi:%v,%v,%v,\n", low, mid, high)
		mid = (low + high) / 2
		//fmt.Printf("mid is %v:", mid)
		if target < from[mid] {
			high = mid - 1
		} else if target > from[mid] {
			low = mid + 1
		} else {
			return mid
		}

	}
	return -1
}
func InsertSearch(from []int, target int) int {
	var low, high, mid int
	low = 0
	high = len(from) - 1
	for low <= high {

		mid = low + (target-from[low])*(high-low)/(from[high]-from[low]) //下标最大为high-1
		//fmt.Printf("riaoshi:%v,%v,%v,\n", low, mid, high)

		if target < from[mid] {
			high = mid - 1
		} else if target > from[mid] {
			low = mid + 1
		} else {
			return mid
		}

	}
	return -1
}
func getFin(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return getFin(n-1) + getFin(n-2)
}
func getFinSlice(l int) []int {
	s := make([]int, l)
	for i := 0; i < l; i++ {
		s[i] = getFin(i)
	}
	return s
}
func FinSearch(from []int, target int) int {
	var low, high, mid, k int
	var F []int

	low = 0
	high = len(from) - 1
	F = getFinSlice(high)
	k = 0
	//fmt.Println(F)
	for high > F[k]-1 {
		k++
	}
	//fmt.Println(k)
	for i := high; i < F[k]-1; i++ {
		from = append(from, from[high])
	}
	//fmt.Println(form)
	for low <= high {
		mid = low + F[k-1] - 1
		if target < from[mid] {
			high = mid - 1
			k = k - 1
		} else if target > from[mid] {
			low = mid + 1
			k = k - 2
		} else {
			if mid <= high {
				return mid
			} else {
				return high
			}
		}

	}
	return -1
}

//二叉排序树搜索f只是一个暂存功能存储未查找到情况下的父节点
func (t *BSTTree) SearchBST(key int, f *BSTTree) (p *BSTTree, result bool) {
	if t == nil {
		p = f
		return p, false
	} else if key == t.Data {
		p = t
		return p, true
	} else if key < t.Data {
		return t.Leftchild.SearchBST(key, t)
	} else {
		return t.Rightchild.SearchBST(key, t)
	}
}

//二叉排序树
func (t *BSTTree) InsertBST(key int) bool {
	s := &BSTTree{}

	if p, ok := t.SearchBST(key, nil); ok == false {
		s.Data = key
		if p == nil {
			fmt.Println("run")
			t = s
		} else if key < p.Data {
			p.Leftchild = s
		} else {
			p.Rightchild = s
		}
		return true
	} else {
		return false
	}

}
func (t *BSTTree) DeleteBST(key int) bool {
	if t == nil {
		return false
	} else if key == t.Data {
		return t.Delete()
	} else if key < t.Data {
		return t.Leftchild.DeleteBST(key)
	} else {
		return t.Rightchild.DeleteBST(key)
	}
}
func (t *BSTTree) Delete() bool {
	var parent, s *BSTTree
	if t.Rightchild == nil {
		//右子树为空
		t = t.Leftchild
	} else if t.Leftchild == nil {
		//左子树为空
		t = t.Rightchild
	} else {
		//左右均不为空
		parent = t
		s = t.Leftchild
		for s.Rightchild != nil {
			parent = s
			s = s.Rightchild
		}
		t.Data = s.Data
		if parent != t {
			parent.Rightchild = s.Leftchild

		} else {
			parent.Leftchild = s.Leftchild
		}
	}
	return true
}
