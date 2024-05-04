package main

import (
	"context"
	"fmt"
	vendor "github.com/LucasGois1/src/domain"
	"github.com/LucasGois1/src/domain/customer"
	"github.com/LucasGois1/src/domain/flavor"
	"github.com/LucasGois1/src/domain/tasks"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg = sync.WaitGroup{}
	c := context.WithValue(context.Background(), "wg", &wg)
	ctx, _ := context.WithDeadline(c, time.Now().Add(20*time.Second))

	aCustomer, aCustomer2, aCustomer3, aCustomer4, aCustomer5 := executeUseCase(ctx)

	aCustomer.EatIceCream()
	aCustomer2.EatIceCream()
	aCustomer3.EatIceCream()
	aCustomer4.EatIceCream()
	aCustomer5.EatIceCream()

	fmt.Println("Time elapsed: ", time.Since(start))
	fmt.Println("End of the program!")
}

func executeUseCase(ctx context.Context) (customer.Customer, customer.Customer, customer.Customer, customer.Customer, customer.Customer) {
	aCustomer := customer.Customer{Name: "Beatriz", FavoriteFlavor: flavor.Pistachio}
	aVendor := vendor.Vendor{Name: "Lucas"}

	order1 := tasks.AttendanceQueue{
		Vendor:   &aVendor,
		Customer: &aCustomer,
	}

	aCustomer2 := customer.Customer{Name: "Monica", FavoriteFlavor: flavor.Lemon}
	aVendor2 := vendor.Vendor{Name: "Marcio"}

	order2 := tasks.AttendanceQueue{
		Vendor:   &aVendor2,
		Customer: &aCustomer2,
	}

	aCustomer3 := customer.Customer{Name: "Patricia", FavoriteFlavor: flavor.Strawberry}
	aVendor3 := vendor.Vendor{Name: "Fernando"}

	order3 := tasks.AttendanceQueue{
		Vendor:   &aVendor3,
		Customer: &aCustomer3,
	}

	aCustomer4 := customer.Customer{Name: "Maria", FavoriteFlavor: flavor.Yogurt}
	aVendor4 := vendor.Vendor{Name: "Joao"}

	order4 := tasks.AttendanceQueue{
		Vendor:   &aVendor4,
		Customer: &aCustomer4,
	}

	aCustomer5 := customer.Customer{Name: "Julieta", FavoriteFlavor: flavor.Chocolate}
	aVendor5 := vendor.Vendor{Name: "Romeu"}

	order5 := tasks.AttendanceQueue{
		Vendor:   &aVendor5,
		Customer: &aCustomer5,
	}

	allOrders := tasks.NewAllOrders()

	allOrders.Enqueue(&order1)
	allOrders.Enqueue(&order2)
	allOrders.Enqueue(&order3)
	allOrders.Enqueue(&order4)
	allOrders.Enqueue(&order5)

	iceCreams := allOrders.ExecuteAll(ctx)

	aCustomer.HoldIceCream(iceCreams[0])
	aCustomer2.HoldIceCream(iceCreams[1])
	aCustomer3.HoldIceCream(iceCreams[2])
	aCustomer4.HoldIceCream(iceCreams[3])
	aCustomer5.HoldIceCream(iceCreams[4])

	return aCustomer, aCustomer2, aCustomer3, aCustomer4, aCustomer5
}
