package observer

import "fmt"

// ISubject subject
type ISubject interface {
	Register(observer IObsever)
	Remove(observer IObsever)
	Notify(observer IObsever)
}

// IObsever 观察者
type IObsever interface {
	Update(msg string)
}

// Subject Subject
type Subject struct {
	observers []IObsever
}

// Register 注册
func (sub *Subject) Register(observer IObsever) {
	sub.observers = append(sub.observers, observer)
}

// Remove 移除观察者
func (sub *Subject) Remove(observer IObsever) {
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
	}
}

// Notify 通知
func (sub *Subject) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}

// Obsever1 Obsever1
type Obsever1 struct{}

// Update 实现观察者接口
func (Obsever1) Update(msg string) {
	fmt.Printf("Obsever1: %s", msg)
}

// Obsever2 Obsever2
type Obsever2 struct{}

// Update 实现观察者接口
func (Obsever2) Update(msg string) {
	fmt.Printf("Obsever2: %s", msg)
}
