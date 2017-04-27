package design

import (
	"errors"
)

type Operation interface {
	GetResult() (result float64)
}
type Add struct {
	Data []float64
}
type Muti struct {
	Data []float64
}
type Min struct {
	Data []float64
}
type Divi struct {
	Data []float64
}

func GetOperation(what string, data ...float64) Operation {
	switch what {
	case "add":
		return Add{Data: data}
	case "min":
		return Min{Data: data}
	case "muti":
		return Muti{Data: data}
	case "divi":
		return Divi{Data: data}
	default:
		return Add{Data: []float64{1, 2, 3, 5}}

	}
}
func (m Muti) GetResult() (result float64) {
	result = 1
	for _, v := range m.Data {
		result = result * v
	}
	return
}
func (a Add) GetResult() (result float64) {
	for _, v := range a.Data {
		result = result + v
	}
	return
}
func (m Min) GetResult() (result float64) {
	result = m.Data[0]
	for i := 1; i < len(m.Data); i++ {
		result = result - m.Data[i]
	}
	return
}
func (d Divi) GetResult() (result float64) {
	result = d.Data[0]
	for i := 1; i < len(d.Data); i++ {
		result = result / d.Data[i]
	}
	return
}

//
var operationfactory = make(map[string]func(data ...float64) Operation)
var oper = []string{"+", "-", "*"}

func NewAdd(data ...float64) Operation {
	return Add{Data: data}
}
func NewMin(data ...float64) Operation {
	return Min{Data: data}
}
func NewMuti(data ...float64) Operation {
	return Muti{Data: data}
}
func NewDivi(data ...float64) Operation {
	return Divi{Data: data}
}
func NewOperation(o string, data ...float64) Operation {
	_, ok := operationfactory[o]
	if ok {
		return operationfactory[o](data...)
	} else {
		panic(errors.New("该方法不存在"))
	}
}

//可以达到动态注册的目的
func RegisterFactory(o string, f func(data ...float64) Operation) {
	operationfactory[o] = f
}
func init() {
	RegisterFactory("+", NewAdd)
	RegisterFactory("-", NewMin)
	RegisterFactory("*", NewMuti)
}
