package learn

import (
	"fmt"

	//"fmt"
	"testing"
)

func Test_getFin(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log("获取数组如下")
		//fmt.Println(getFin(i))
	}

}
func Test_getFinSlice(t *testing.T) {

	s := getFinSlice(5)
	correct := []int{0, 1, 1, 2, 3}
	for i := 0; i < 5; i++ {
		if s[i] != correct[i] {
			t.Error("获取数组错误")
		}
	}
}
func Test_BinarySearch(t *testing.T) {
	from := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		if i != BinarySearch(from, i) {
			t.Error("二分查找错误")
		}
	}

}
func Benchmark_BinarySearch(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	from := []int{3, 5, 6, 8, 10, 12}
	BinarySearch(from, 8)

}
func Test_InsertSearch(t *testing.T) {
	from := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		/*fmt.Println("result")
		fmt.Printf("result:%v", is)
		fmt.Println("------result")*/
		if i != InsertSearch(from, i) {
			t.Error("插值查找错误")
		}

	}
}
func Test_FinSearch(t *testing.T) {
	from := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	fmt.Println(FinSearch(from, 9))
}

//艰难的测试通过了不知道还会有什么问题
func Test_BSTTREE(t *testing.T) {
	a := []int{62, 88, 58, 47, 35, 73, 51, 99, 37, 93}
	T := &BSTTree{Data: 62}
	/*l := &BSTTree{Data: 58}
	r := &BSTTree{Data: 88}
	rr := &BSTTree{Data: 99}
	r.Rightchild = rr
	T.Leftchild = l
	T.Rightchild = r
	fmt.Println(T.SearchBST(62, nil))
	fmt.Println(T.SearchBST(58, nil))
	fmt.Println(T.SearchBST(88, nil))
	fmt.Println(T.SearchBST(99, nil))
	fmt.Println(T.SearchBST(100, nil))*/
	for i := 0; i < 10; i++ {
		T.InsertBST(a[i])
	}
	T.PreOrderTraverse()
	T.InsertBST(100)
	T.InsertBST(1)
	T.PreOrderTraverse()
	/*T.InOrderTraverse()
	T.AfterOrderTraverse()*/
}
