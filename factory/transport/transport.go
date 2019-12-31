package main

type Transport interface {
	deliver()
	packing()
}

type TransportationFactory interface {
	Create() Transport
}

type Factory struct {
	Factories []TransportationFactory
}

func (f *Factory) CreateTransportMethods() {
	for _, f := range f.Factories {
		trans := f.Create()
		trans.packing()
		trans.deliver()
	}
}
