package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	sub.Register(&Obsever1{})
	sub.Register(&Obsever2{})
	sub.Notify("hi")
}
