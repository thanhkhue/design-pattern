package main

import "fmt"

type Drone struct{}

type DroneFactory struct{}

func (factory DroneFactory) Create() Transport {
	return &Drone{}
}

func (d *Drone) deliver() {
	fmt.Println("Deliver by Drone")
}

func (d *Drone) packing() {
	fmt.Println("Packing product into a box")
}
