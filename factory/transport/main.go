package main



func main() {
	f := &Factory{
		Factories: []TransportationFactory{
			DroneFactory{},
			ShipFactory{},
			TruckFactory{},
		},
	}
	
	f.CreateTransportMethods()
}
