package main

import (
	"fmt"
	vendor "github.com/LucasGois1/src/domain"
	"github.com/LucasGois1/src/domain/flavor"
	"time"
)

func main() {
	start := time.Now()
	aVendor := vendor.Vendor{Name: "Lucas"}

	aVendor.MakeIceCream(flavor.Chocolate)
	fmt.Println("Done! Enjoy your Chocolate ice cream!")

	aVendor.MakeIceCream(flavor.Vanilla)
	fmt.Println("Done! Enjoy your Vanilla ice cream!")

	aVendor.MakeIceCream(flavor.Strawberry)
	fmt.Println("Done! Enjoy your Strawberry ice cream!")

	aVendor.MakeIceCream(flavor.Mint)
	fmt.Println("Done! Enjoy your Mint ice cream!")

	aVendor.MakeIceCream(flavor.Grape)
	fmt.Println("Done! Enjoy your Grape ice cream!")

	aVendor.MakeIceCream(flavor.Lemon)
	fmt.Println("Done! Enjoy your Lemon ice cream!")

	aVendor.MakeIceCream(flavor.Yogurt)
	fmt.Println("Done! Enjoy your Yogurt ice cream!")

	fmt.Println("Time elapsed: ", time.Since(start))
}
