package design

import (
	"testing"
)

func Test_mem(t *testing.T) {
	var o Origin = &Originator{State: "on"}
	if o.Show() != "on" {
		t.Error("状态不对")
	}
	c := Caretaker{m: o.CreateMemento()}
	o.Set("off")
	if o.Show() != "off" {
		t.Error("状态不对 未改变")
	}
	o.SetMemento(c.m)
	if o.Show() != "on" {
		t.Error("状态错误,未成功还原")
	}
}
