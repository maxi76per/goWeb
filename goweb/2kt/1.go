package main

import "fmt"

// Подсистема 1
type SubsystemA struct{}

func (s *SubsystemA) OperationA() {
	fmt.Println("Subsystem A: Operation A")
}

// Подсистема 2
type SubsystemB struct{}

func (s *SubsystemB) OperationB() {
	fmt.Println("Subsystem B: Operation B")
}

// Фасад
type Facade struct {
	a *SubsystemA
	b *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		a: &SubsystemA{},
		b: &SubsystemB{},
	}
}

func (f *Facade) Operation() {
	f.a.OperationA()
	f.b.OperationB()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}