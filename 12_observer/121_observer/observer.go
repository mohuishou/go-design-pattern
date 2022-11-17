package observer

import "fmt"

// ISubject subject
type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

// IObserver 观察者
type IObserver interface {
	Update(msg string)
	GetName() string
}

// Subject Subject
type Subject struct {
	observers []IObserver
}

// Register 注册
func (sub *Subject) Register(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}

// Remove 移除观察者
func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.observers {
		if ob.GetName() == observer.GetName() {
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

// Observer1 Observer1
type Observer1 struct {
	Name string
}

// Update 实现观察者接口
func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}

func (o Observer1) GetName() string {
	return o.Name
}

// Observer2 Observer2
type Observer2 struct {
	Name string
}

// Update 实现观察者接口
func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s", msg)
}

func (o Observer2) GetName() string {
	return o.Name
}
