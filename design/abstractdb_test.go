package design

import (
	"testing"
)

func Test_abstractdb(t *testing.T) {
	f := SqlFactory{}
	u := f.CreateUser()
	d := f.CreateDepartment()
	u.Get()
	u.Insert(u)
	d.Get()
	d.Insert(d)

}
