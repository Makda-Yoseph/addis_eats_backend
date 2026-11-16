package order

import (
	"errors"
	"fmt"
	"math"

	"github.com/Makda-Yoseph/addis_eats_backend.git/internal/menu"
	"github.com/Makda-Yoseph/addis_eats_backend.git/internal/delivery"
)

var ErrEmptyOrder = errors.New("order has no items!")
const FreeDeliveryThresholdSantim = 50000


type Order struct{
	Items []menu.Dish
	Area string
}
func (o *Order) AddItem(d menu.Dish){
	o.Items = append(o.Items, d)
}

func(o Order) SubtotalSantim()int{
	total := 0
	for  _, dish := range o.Items{
		total += int(math.Round(dish.PriceSantim*100))
	}
	return total
}

func(o Order) Total()(int,error){
	if len(o.Items)==0{
		return 0,ErrEmptyOrder
	}
	subtotal := o.SubtotalSantim()
	serviceFee := subtotal* 5 /1000
	deliveryFee,err := delivery.FeeSantim(o.Area)
	if err!= nil{
		return 0,fmt.Errorf("calculate total: %w", err)
	}
	if subtotal > FreeDeliveryThresholdSantim{
		deliveryFee = 0
	}
	return deliveryFee + subtotal + serviceFee , nil

}

