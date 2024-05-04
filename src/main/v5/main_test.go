package main

import (
	"context"
	"github.com/LucasGois1/src/domain/customer"
	"github.com/LucasGois1/src/domain/flavor"
	"github.com/LucasGois1/src/domain/ice_cream"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_executeUseCase(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name                   string
		args                   args
		expectedFirstCustomer  customer.Customer
		expectedSecondCustomer customer.Customer
		expectedThirdCustomer  customer.Customer
		expectedFourthCustomer customer.Customer
		expectedFifthCustomer  customer.Customer
	}{
		{
			name:                   "Should return the customers with correct ice creams that each one ordered",
			args:                   args{ctx: context.WithValue(context.Background(), "wg", &sync.WaitGroup{})},
			expectedFirstCustomer:  customer.Customer{Name: "Beatriz", FavoriteFlavor: 9, IceCream: &ice_cream.IceCream{Flavor: flavor.Pistachio, WhoDidIt: "Lucas"}},
			expectedSecondCustomer: customer.Customer{Name: "Monica", FavoriteFlavor: 3, IceCream: &ice_cream.IceCream{Flavor: flavor.Chocolate, WhoDidIt: "Marcio"}},
			expectedThirdCustomer:  customer.Customer{Name: "Patricia", FavoriteFlavor: 6, IceCream: &ice_cream.IceCream{Flavor: flavor.Yogurt, WhoDidIt: "Fernando"}},
			expectedFourthCustomer: customer.Customer{Name: "Maria", FavoriteFlavor: 7, IceCream: &ice_cream.IceCream{Flavor: flavor.Strawberry, WhoDidIt: "Joao"}},
			expectedFifthCustomer:  customer.Customer{Name: "Julieta", FavoriteFlavor: 8, IceCream: &ice_cream.IceCream{Flavor: flavor.Vanilla, WhoDidIt: "Romeu"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customer1, customer2, customer3, customer4, customer5 := executeUseCase(tt.args.ctx)
			assert.ObjectsAreEqualValues(tt.expectedFirstCustomer, customer1)
			assert.ObjectsAreEqualValues(tt.expectedSecondCustomer, customer2)
			assert.ObjectsAreEqualValues(tt.expectedThirdCustomer, customer3)
			assert.ObjectsAreEqualValues(tt.expectedFourthCustomer, customer4)
			assert.ObjectsAreEqualValues(tt.expectedFifthCustomer, customer5)
		})
	}
}
