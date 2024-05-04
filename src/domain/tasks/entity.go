package tasks

import (
	"context"
	"fmt"
	vendor "github.com/LucasGois1/src/domain"
	"github.com/LucasGois1/src/domain/customer"
	"github.com/LucasGois1/src/domain/ice_cream"
	"sync"
)

type AttendanceQueue struct {
	Vendor   *vendor.Vendor
	Customer *customer.Customer
}

type VendorBell struct {
	Vendor *vendor.Vendor
}

func (a *AttendanceQueue) Attend(ringBell chan *VendorBell) {
	a.Vendor.MakeIceCream(a.Customer.FavoriteFlavor)

	fmt.Println(a.Vendor.Name, "says:", "Done! Enjoy your ice cream", a.Customer.Name, "!")
	ringBell <- &VendorBell{
		Vendor: a.Vendor,
	}

}

type AllOrders struct {
	TaskAndFinishedBell map[*AttendanceQueue]chan *VendorBell
	AttendanceOrder     map[*vendor.Vendor]int
}

func NewAllOrders() *AllOrders {
	return &AllOrders{
		TaskAndFinishedBell: make(map[*AttendanceQueue]chan *VendorBell),
		AttendanceOrder:     make(map[*vendor.Vendor]int),
	}
}

func (a *AllOrders) Enqueue(task *AttendanceQueue) {
	a.TaskAndFinishedBell[task] = make(chan *VendorBell)
	a.AttendanceOrder[task.Vendor] = len(a.AttendanceOrder)
}

func (a *AllOrders) ExecuteAll(ctx context.Context) []*ice_cream.IceCream {
	go func() {
		for task, result := range a.TaskAndFinishedBell {
			go func(task *AttendanceQueue, finishedBell chan *VendorBell) {
				task.Attend(finishedBell)
				fmt.Println("Vendor", task.Vendor.Name, "finished the task!")
			}(task, result)
		}
	}()

	fmt.Println("Waiting for all vendors to finish their tasks...")
	wg := ctx.Value("wg").(*sync.WaitGroup)

	var finishedTasks []*VendorBell
	for task, result := range a.TaskAndFinishedBell {
		wg.Add(1)
		go func() {
			finishedTasks = append(finishedTasks, <-result)
			fmt.Println("Vendor", task.Vendor.Name, "ice cream has been collected!")
			wg.Done()
		}()
	}

	wg.Wait()
	iceCreams := make([]*ice_cream.IceCream, len(finishedTasks))
	for _, finishedTask := range finishedTasks {
		iceCreams[a.AttendanceOrder[finishedTask.Vendor]] = finishedTask.Vendor.IceCream
	}

	return iceCreams
}
