package design

import (
	"testing"
)

func Test_composite(t *testing.T) {
	cmap := make(map[string]Company)
	var root Company = &ConcreteCompany{Name: "总公司", children: cmap}
	root.Add(&HRDepartment{Name: "总公司人力资源"})
	var comp Company = &ConcreteCompany{Name: "华东分公司", children: make(map[string]Company)}
	comp.Add(&HRDepartment{Name: "华东分公司人力资源"})
	var banshi Company = &ConcreteCompany{Name: "华东浙江办事处", children: make(map[string]Company)}
	comp.Add(banshi)
	root.Add(comp)
	root.LineOfDuty()
	root.Display(2)
}
