package design

import (
	"testing"
)

func Test_fly(t *testing.T) {
	f := FlyWeightFactory("X")
	f.Operation(1)
	u := FlyWeightFactory("U")
	u.Operation(3)
}
