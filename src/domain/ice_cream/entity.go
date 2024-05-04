package ice_cream

import (
	"fmt"
	"github.com/LucasGois1/src/domain/flavor"
)

type IceCream struct {
	Flavor   flavor.Flavor
	Cost     float64
	WhoDidIt string
}

func (i *IceCream) String() string {
	return fmt.Sprintf("Flavor: %s, Cost: %.2f, Who did it: %s", flavor.Name(i.Flavor), i.Cost, i.WhoDidIt)
}
