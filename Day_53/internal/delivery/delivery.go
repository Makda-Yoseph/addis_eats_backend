package delivery

import(
	"fmt"
	"errors"
)

var ErrUnservedArea = errors.New("Unserved Area!")

var areaFee = map[string]int{
	"Bole":500,
	"Gerji":200,
	"piassa":400,
}

func FeeSantim(area string) (int, error){
	Fee,ok := areaFee[area]
	if !ok{
		return 0, fmt.Errorf("%w : %q", ErrUnservedArea, area)
	}
	return Fee,nil
}
