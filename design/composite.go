package design

import (
	"fmt"
)

func P(s string, n int) string {
	r := ""
	for i := 0; i < n; i++ {
		r = r + fmt.Sprintf("%v", s)
	}
	return r
}

//将对象 组合成树形结构以表示部分整体的层次
//结构，组合模式使得用户对单个对象和组合对象的使用具有一致性
type Company interface {
	Add(co Company)
	GetName() string
	Remove(co Company)
	Display(depth int)
	LineOfDuty()
}
type ConcreteCompany struct {
	children map[string]Company
	Name     string
}

func (c *ConcreteCompany) Add(co Company) {
	c.children[co.GetName()] = co
}
func (c *ConcreteCompany) Remove(co Company) {
	delete(c.children, co.GetName())
}
func (c ConcreteCompany) GetName() string {
	return c.Name
}
func (c ConcreteCompany) Display(depth int) {
	fmt.Println(P("-", depth), c.Name)
	for _, v := range c.children {
		v.Display(depth + 2)
	}
}
func (c ConcreteCompany) LineOfDuty() {
	for _, v := range c.children {
		v.LineOfDuty()
	}
}

type HRDepartment struct {
	Name string
}

func (h HRDepartment) GetName() string {
	return h.Name
}
func (h *HRDepartment) Add(co Company) {

}
func (h *HRDepartment) Remove(co Company) {

}
func (h HRDepartment) Display(depth int) {
	fmt.Println(P("-", depth), h.Name)
}
func (h HRDepartment) LineOfDuty() {
	fmt.Println(h.Name, "员工招聘培训")
}
