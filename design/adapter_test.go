package design

import (
	"testing"
)

/*func Test_client(t *testing.T) {
	client()
}*/
func Test_adapt(t *testing.T) {
	var p Player = Translator{Name: "姚明"}
	p.Attack()

}
