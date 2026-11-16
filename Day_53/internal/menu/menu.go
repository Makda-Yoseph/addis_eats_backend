package menu


type Dish struct{
	Name string
	PriceSantim float64
	isFasting bool 
}
func menu() []Dish{
	return []Dish{
		{Name:"Doro", PriceSantim: 290.2, isFasting: false},
		{Name:"misir", PriceSantim: 200.2, isFasting: true},
		{Name:"shiro", PriceSantim: 300.0, isFasting: true},
		{Name:"Tibs", PriceSantim: 900.2, isFasting: false},
		{Name:"kitfo", PriceSantim: 400.2, isFasting: false},
		{Name:"kkl", PriceSantim: 892.2, isFasting: false},
		{Name:"Gomen", PriceSantim: 130.2, isFasting: true},
		{Name:"firfir", PriceSantim: 452.2, isFasting: false},

	}

}