package design

import (
	"testing"
)

func Test_bridge(t *testing.T) {
	p := new(Nokia)
	g := Game{}
	p.SetSoft(g)
	p.Run()
}
