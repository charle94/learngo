package design

import (
	"fmt"
)

type Template interface {
	OptionFunction_A()
	OptionFunction_B()
}

//称之为模板方法吧
func TemplateFubction(t Template) {
	t.OptionFunction_A()
	t.OptionFunction_B()
	fmt.Println("templatemethod get")
}

type TemplateInstance struct {
}

func (t TemplateInstance) OptionFunction_A() {

}
func (t TemplateInstance) OptionFunction_B() {

}

//采用如下方法的调用a.templatefunction输出为templatemethod get
//调用不到父结构体的方法，只是使用了匿名字段中的functiona和functionb
/*func (t TemplateInstance) TemplateFubction() {
t.OptionFunction_A()
t.OptionFunction_B()
fmt.Println("templatemethod get")
}*/

type TemplateA struct {
	TemplateInstance
}

func (t TemplateA) OptionFunction_A() {
	fmt.Println("template_a functiona ")
}
func (t TemplateA) OptionFunction_B() {
	fmt.Println("template_a functinob")
}
