package design

import (
	"testing"
)

func Test_command(t *testing.T) {
	r := ConcreteReceiver{}
	b := Bab{}
	var c1 Command = &ConcreteCommand{r: b}
	var c2 Command = &ConcreteCommand{r: r}
	i := Invoker{}
	i.AddCommand(c1)
	i.AddCommand(c2)
	i.ExecuteCommand()
}
