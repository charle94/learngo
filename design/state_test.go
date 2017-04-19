package design

import (
	"testing"
)

/*func Test_state(t *testing.T) {
	c := Context{}
	c.S = StateA{}
	c.Request()
	c.Request()
	c.Request()

}*/
//example test
func Test_state(t *testing.T) {
	w := Work{Finshed: false, S: ForeNoon{}}
	w.Time(9)
	w.Time(13)
	w.Time(17)
	w.Time(19)
	w.Finshed = true
	w.Time(19)

}
