package lab7

type Fruit struct {
	Name      string
	Freshness string
	Price     float64
}

func (f *Fruit) GetName() string {
	return f.Name
}

func (f Fruit) GetPrice() float64 {
	return f.Price
}

func (f *Fruit) SetPrice(price float64) {
	f.Price = price
}

func (m *Fruit) ApplyDiscount(discount float64) {
	f.Price -= f.Price * discount / 100
}