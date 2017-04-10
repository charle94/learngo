package design

import (
	"testing"
)

func Test_factory(t *testing.T) {
	f := UnderGraduateFactory{}
	lei := f.CreateLei()
	lei.Sweep()
	lei.Wash()
	v := VolunteerFactory{}
	lei = v.CreateLei()
	lei.Sweep()
	lei.Wash()
}
