package main

import "fmt"

type Ship struct{}

type ShipFactory struct{}

func (factory ShipFactory) Create() Transport {
	return &Ship{}
}

func (s *Ship) deliver() {
	fmt.Println("Deliver by Ship")
}

func (s *Ship) packing() {
	fmt.Println("Packing product into a container")
}
