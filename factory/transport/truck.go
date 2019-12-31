package main

import "fmt"

type Truck struct{}

type TruckFactory struct{}

func (factory TruckFactory) Create() Transport {
	return &Truck{}
}

func (t *Truck) deliver() {
	fmt.Println("Deliver by Truct")
}

func (t *Truck) packing() {
	fmt.Println("Packing product into a box")
}
