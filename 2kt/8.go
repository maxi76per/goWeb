package main

import "fmt"

// Состояние
type State interface {
	Handle(context *Context)
}

// Конкретное состояние A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("Обработка в состоянии A")
	context.SetState(&ConcreteStateB{})
}

// Конкретное состояние B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("Обработка в состоянии B")
	context.SetState(&ConcreteStateA{})
}

// Контекст
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func main() {
	context := &Context{state: &ConcreteStateA{}}
	context.Request()
	context.Request()
}