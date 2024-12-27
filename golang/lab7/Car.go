package lab7

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

func (c *Car) GetPrice() float64 {
	return c.price
}

func (c *Car) SetPrice(price float64) {
	c.Price = price
}

func (c *Car) ApplyDiscount(discount float64) {
	c.Price -= c.Price * discount / 100
}