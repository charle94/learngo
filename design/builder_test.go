package design

import (
	"testing"
)

func Test_builder(t *testing.T) {
	f := FatPerson{}
	d := PersonDirector{}
	d.Set(f)
	d.Build()
}
