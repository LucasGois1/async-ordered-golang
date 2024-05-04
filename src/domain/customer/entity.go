package customer

import (
	"fmt"
	"github.com/LucasGois1/src/domain/flavor"
	"github.com/LucasGois1/src/domain/ice_cream"
)

type Customer struct {
	Name           string
	FavoriteFlavor flavor.Flavor
	IceCream       *ice_cream.IceCream
}

func (c *Customer) HoldIceCream(iceCream *ice_cream.IceCream) {
	fmt.Println(c.Name, "says:", "Thank you! I'll hold my ice cream!")
	c.IceCream = iceCream
	fmt.Println(c.Name, "says:", "I ordered a", flavor.Name(c.FavoriteFlavor), "ice cream and got", flavor.Name(c.IceCream.Flavor), "!")
}

func (c *Customer) EatIceCream() {
	c.IceCream = nil
	fmt.Println(c.Name, "says:", "Yummy! my ice cream was delicious!")
}

// toString
func (c *Customer) String() string {
	return fmt.Sprintf("Name: %s, Favorite Flavor: %s, Ice Cream: %s", c.Name, flavor.Name(c.FavoriteFlavor), flavor.Name(c.IceCream.Flavor))
}
