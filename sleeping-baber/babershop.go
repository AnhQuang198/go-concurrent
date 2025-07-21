package main

import (
	"github.com/fatih/color"
	"sync"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarber  int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
	TotalClients    int
	mu              sync.Mutex
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarber++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			// if there are no clients, the baber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s take a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close this goroutince
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)

	// increase client count
	shop.mu.Lock()
	shop.TotalClients++
	shop.mu.Unlock() //có thể sử dụng defer nhưng ko cần thiết trong case này vì logic không quan trọng và ko có panic
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day.")

	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarber; a++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Green("---------------------------------------------------------------------")
	color.Green("The barbershop is now closed for the day, and everyone has gone home.")
	color.Green("Total clients served for day: %d", shop.TotalClients)
}

func (shop *BarberShop) addClient(client string) {
	// print out a message
	color.Green("*** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Red("The shop is already closed. so %s leaves!", client)
	}
}
