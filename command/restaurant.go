package main

import "fmt"

// Order command
type Order interface {
	execute()
}

type Chef struct{}

func (c *Chef) MakePizza() {
	fmt.Println("Making Pizza")
}

func (c *Chef) MakeSteak() {
	fmt.Println("Making Steak")
}

// PizzaChef concrete command of making pizza
type PizzaChef struct {
	Chef *Chef
}

func (p *PizzaChef) execute() {
	p.Chef.MakePizza()
}

// SteakChef concrete command of making steak
type SteakChef struct {
	Chef *Chef
}

func (s *SteakChef) execute() {
	s.Chef.MakeSteak()
}

// Waitress
type Waitress struct {
	orders []Order
}

func (w *Waitress) PlaceOrder(order Order) {
	order.execute()
}

func main() {
	// receiver
	c := new(Chef)

	// concrete commands
	makePizza := &PizzaChef{Chef: c}
	makeSteak := &SteakChef{Chef: c}

	// invoker
	w := new(Waitress)
	w.PlaceOrder(makePizza)
	w.PlaceOrder(makeSteak)

}
