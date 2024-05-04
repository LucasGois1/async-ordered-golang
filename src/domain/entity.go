package vendor

import (
	"fmt"
	"github.com/LucasGois1/src/domain/flavor"
	"github.com/LucasGois1/src/domain/ice_cream"
	"time"
)

type Vendor struct {
	Name     string
	IceCream *ice_cream.IceCream
}

func (v *Vendor) MakeIceCream(flavor flavor.Flavor) {
	fmt.Println(v.Name, "says:", "Making an ice cream...")
	time.Sleep(time.Duration(flavor) * time.Second)

	v.IceCream = &ice_cream.IceCream{
		Flavor:   flavor,
		WhoDidIt: v.Name,
	}
}

// toString
func (v *Vendor) String() string {
	return fmt.Sprintf("Name: %s, Ice Cream: %s", v.Name, v.IceCream.String())
}
