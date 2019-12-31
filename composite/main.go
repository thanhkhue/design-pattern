package main

// StoreItem store's item
type StoreItem interface {
	// calculate total price
	getPrice() int
}

type Packages struct {
	Items []StoreItem
}

// Product store's product
type Product struct {
	Name  string
	Price int
}

// Box contains smaller boxex and products
type Box struct {
	SmallerBoxes []*Box
	Products     []*Product
}

func (p *Product) getPrice() int {
	return p.Price
}

func (b *Box) getPrice() int {
	price := 0
	if len(b.SmallerBoxes) > 0 {
		for i := 0; i < len(b.SmallerBoxes); i++ {
			price += b.SmallerBoxes[i].getPrice()
		}
	}

	if len(b.Products) > 0 {
		for i := 0; i < len(b.Products); i++ {
			price += b.Products[i].getPrice()
		}
	}

	return price
}

func (p *Packages) getPrice() int {
	price := 0
	for i := 0; i < len(p.Items); i++ {
		price += p.Items[i].getPrice()
	}
	return price
}

func main() {

	for i := 0; i < 4; i++ {

	}
}
