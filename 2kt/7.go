package main

import "fmt"

// Стратегия
type Strategy interface {
	Execute(a, b int) int
}

// Конкретная стратегия A
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute(a, b int) int {
	return a + b
}

// Конкретная стратегия B
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute(a, b int) int {
	return a - b
}

// Контекст
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	context := &Context{}
	context.SetStrategy(&ConcreteStrategyA{})
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))

	context.SetStrategy(&ConcreteStrategyB{})
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5))
}