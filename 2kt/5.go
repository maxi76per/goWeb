package main

import "fmt"

// Обработчик
type Handler interface {
	SetNext(handler Handler)
	Handle(request int)
}

// Базовый обработчик
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) Handle(request int) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

// Конкретный обработчик A
type ConcreteHandlerA struct {
	BaseHandler
}

func (c *ConcreteHandlerA) Handle(request int) {
	if request < 10 {
		fmt.Println("ConcreteHandlerA обработал запрос")
	} else {
		c.BaseHandler.Handle(request)
	}
}

// Конкретный обработчик B
type ConcreteHandlerB struct {
	BaseHandler
}

func (c *ConcreteHandlerB) Handle(request int) {
	if request >= 10 && request < 20 {
		fmt.Println("ConcreteHandlerB обработал запрос")
	} else {
		c.BaseHandler.Handle(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerA.SetNext(handlerB)

	requests := []int{5, 15, 25}
	for _, req := range requests {
		handlerA.Handle(req)
	}
}