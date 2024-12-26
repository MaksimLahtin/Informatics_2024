package lab7

import "fmt"

type Fruit struct {
	Name      string
	Freshness string
	Price     float64
}

func (f *Fruit) GetName() string {
	return m.Name
}

func (f Fruit) GetPrice() float64 {
	return f.Price
}

func (f *Fruit) SetPrice(price float64) {
	m.Price = price
}

func (m *Fruit) ApplyDiscount(discount float64) {
	m.Price -= m.Price * discount / 100
}