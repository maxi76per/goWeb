package main

import "fmt"

// Продукт
type Product interface {
	Use()
}

// Конкретный продукт A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Используется ConcreteProductA")
}

// Конкретный продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Используется ConcreteProductB")
}

// Создатель
type Creator interface {
	FactoryMethod() Product
}

// Конкретный создатель A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() Product {
	return &ConcreteProductA{}
}

// Конкретный создатель B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() Product {
	return &ConcreteProductB{}
}

func main() {
	creators := []Creator{&ConcreteCreatorA{}, &ConcreteCreatorB{}}
	for _, creator := range creators {
		product := creator.FactoryMethod()
		product.Use()
	}
}