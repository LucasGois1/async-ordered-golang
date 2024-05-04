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
	go func() {
		aVendor.MakeIceCream(flavor.Chocolate)
		fmt.Println("Done! Enjoy your Chocolate ice cream!")
	}()

	go func() {
		aVendor.MakeIceCream(flavor.Vanilla)
		fmt.Println("Done! Enjoy your Vanilla ice cream!")
	}()

	go func() {
		aVendor.MakeIceCream(flavor.Strawberry)
		fmt.Println("Done! Enjoy your Strawberry ice cream!")
	}()

	go func() {
		aVendor.MakeIceCream(flavor.Mint)
		fmt.Println("Done! Enjoy your Mint ice cream!")
	}()

	go func() {
		aVendor.MakeIceCream(flavor.Grape)
		fmt.Println("Done! Enjoy your Grape ice cream!")
	}()

	go func() {
		aVendor.MakeIceCream(flavor.Lemon)
		fmt.Println("Done! Enjoy your Lemon ice cream!")
	}()

	go func() {
		go aVendor.MakeIceCream(flavor.Yogurt)
		fmt.Println("Done! Enjoy your Yogurt ice cream!")
	}()

	time.Sleep(20 * time.Second)

	fmt.Println("Time elapsed: ", time.Since(start))
	fmt.Println("End of the program!")
}
