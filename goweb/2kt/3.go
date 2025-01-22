package main

import "fmt"

// Элемент
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// Конкретный элемент B
type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Посетитель
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Visited ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Visited ConcreteElementB")
}

func main() {
	elements := []Element{&ConcreteElementA{}, &ConcreteElementB{}}
	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}