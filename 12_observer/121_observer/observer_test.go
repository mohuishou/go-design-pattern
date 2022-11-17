package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	sub.Register(&Observer1{})
	sub.Register(&Observer2{})
	sub.Notify("hi")
}

func TestSubject_Remove(t *testing.T) {
	obs1 := &Observer1{
		Name: "obs1",
	}
	obs2 := &Observer2{
		Name: "obs2",
	}

	sub := &Subject{}
	sub.Register(obs1)
	sub.Register(obs2)
	sub.Notify("hi")
	sub.Remove(obs2)
	sub.Notify("hi")
}
