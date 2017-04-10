package design

import (
	"testing"
)

func Test_proxy(t *testing.T) {
	boy := Pursuit{Girl: "mylove"}
	pr := Proxy{Boy: boy}
	pr.GiveDolls()
	pr.GiveFlowers()
}
