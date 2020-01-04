package main

import "fmt"

// Light receiver
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Turn-on the light")
}

func (l *Light) TurnOff() {
	fmt.Println("Turn-off the light")
}

// Command command interface
type Command interface {
	execute()
}

//SwitchOnCommand - ConcreteCommand(The Command for turning on the light)
type SwitchOnCommand struct {
	Light *Light
}

func (cmd *SwitchOnCommand) execute() {
	cmd.Light.TurnOn()
}

//SwitchOffCommand - ConcreteCommand(The Command for turning off the light)
type SwitchOffCommand struct {
	Light *Light
}

func (cmd *SwitchOffCommand) execute() {
	cmd.Light.TurnOff()
}

//Switch - The invoker struct
type Switch struct {
	commands []Command
}

// StoreAndExecute store and execute command
func (sw *Switch) StoreAndExecute(cmd Command) {
	sw.commands = append(sw.commands, cmd)
	cmd.execute()
}

func init() {
	// receiver
	lamp := new(Light)

	// ConcreteCommand
	swOn := &SwitchOnCommand{Light: lamp}
	swOff := &SwitchOffCommand{Light: lamp}

	// Invoker
	sw := new(Switch)
	sw.StoreAndExecute(swOn)
	sw.StoreAndExecute(swOff)
}
