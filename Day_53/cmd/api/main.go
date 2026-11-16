package main

import (
	"errors"
	"fmt"

	"github.com/Makda-Yoseph/addis_eats_backend.git/internal/delivery"
	"github.com/Makda-Yoseph/addis_eats_backend.git/internal/menu"
	"github.com/Makda-Yoseph/addis_eats_backend.git/internal/order"
)

func main(){
	order1 := order.Order{Area: "Bole"}
	order1.AddItem(menu.Dish{Name: "DoroWot", PriceSantim: 200020})
	order1.AddItem(menu.Dish{Name: "pasta", PriceSantim: 30050})

	order2 := order.Order{Area: "Bole"}

	order3 := order.Order{Area: "Ayat"}
	order3.AddItem(menu.Dish{Name: "DoroWot", PriceSantim: 200020})
	order3.AddItem(menu.Dish{Name: "pasta", PriceSantim: 30050})

	PrintOrderResult("order-1", order1)
	PrintOrderResult("order-2", order2)
	PrintOrderResult("order-3", order3)

}

func PrintOrderResult( label string, o order.Order ) {
	total, err := o.Total()
	if err != nil{
		if errors.Is(err, order.ErrEmptyOrder){
			fmt.Printf("%s -> Rejected: your order is empty, please add a dish.\n", label)
			return
		}
		if errors.Is(err, delivery.ErrUnservedArea){
			fmt.Printf("%s -> Rejected: we do not deliver to %q yet.\n", label, o.Area)
			return
		}
		fmt.Printf("%s -> Unexpected error: %v\n", label, err)
		return
	}
	fmt.Printf("%s -> Success! Total to pay: %d Santim\n", label, total)
}