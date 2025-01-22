package main

import "fmt"

// Продукт
type Product struct {
	partA, partB, partC string
}

// Строитель
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

// Конкретный строитель
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &Product{}}
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.partA = "Part A"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.partB = "Part B"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.partC = "Part C"
}

func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

// Директор
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetProduct()
	fmt.Printf("Product: %+v\n", product)
}