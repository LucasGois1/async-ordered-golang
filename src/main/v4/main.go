package main

import (
	"fmt"
	vendor "github.com/LucasGois1/src/domain"
	"github.com/LucasGois1/src/domain/flavor"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg = sync.WaitGroup{}

	aVendor := vendor.Vendor{Name: "Lucas"}
	wg.Add(1)
	go func() {
		aVendor.MakeIceCream(flavor.Chocolate)
		fmt.Println("Done! Enjoy your Chocolate ice cream!")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		aVendor.MakeIceCream(flavor.Vanilla)
		fmt.Println("Done! Enjoy your Vanilla ice cream!")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		aVendor.MakeIceCream(flavor.Strawberry)
		fmt.Println("Done! Enjoy your Strawberry ice cream!")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		aVendor.MakeIceCream(flavor.Mint)
		fmt.Println("Done! Enjoy your Mint ice cream!")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		aVendor.MakeIceCream(flavor.Grape)
		fmt.Println("Done! Enjoy your Grape ice cream!")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		aVendor.MakeIceCream(flavor.Lemon)
		fmt.Println("Done! Enjoy your Lemon ice cream!")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		go aVendor.MakeIceCream(flavor.Yogurt)
		fmt.Println("Done! Enjoy your Yogurt ice cream!")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Time elapsed: ", time.Since(start))
	fmt.Println("End of the program!")
}
