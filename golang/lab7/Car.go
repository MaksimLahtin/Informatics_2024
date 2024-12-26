package laba7

type Car struct {
	price float64
	name  string
	color string
	model string
	brand string
}

func (c *Car) GetName() string {
	return c.name
}

func (c *Car) GetPrice() string {
	return c.price
}

func (c *Car) SetPrice(price float64) {
	m.Price = price
}

func (c *Car) ApplyDiscount(discount float64) {
	m.Price -= m.Price * discount / 100
}