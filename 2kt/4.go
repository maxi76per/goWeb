package main

import "fmt"

// Команда
type Command interface {
	Execute()
}

// Конкретная команда
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Получатель
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver: выполнение действия")
}

// Инициатор
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := NewConcreteCommand(receiver)
	invoker := &Invoker{}
	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}