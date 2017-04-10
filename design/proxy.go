package design

import (
	"fmt"
)

/**
 * example for proxy pattern
 */
/*type Subject interface {
	Request()
}
type RealSubject struct {
}

func (r RealSubject) Request() {

}

type Proxy struct {
	Re Subject
}

func (p Proxy) Request() {
	p.Re.Request()
}*/

type Gift interface {
	GiveDolls()
	GiveFlowers()
}
type Pursuit struct {
	Girl string
}

func (p Pursuit) GiveDolls() {
	fmt.Println("追求者洋娃娃")
}
func (p Pursuit) GiveFlowers() {
	fmt.Println("追求者送花")
}

type Proxy struct {
	Boy Pursuit
}

func (p Proxy) GiveDolls() {
	p.Boy.GiveDolls()
}
func (p Proxy) GiveFlowers() {
	p.Boy.GiveFlowers()
}
