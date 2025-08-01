// This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem
// which illustrates the complexities that arise when there are multiple operating system processes. Here, we have
// a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is
// open, and clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the
// waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to
// sleep until a new client arrives. So the rules are as follows:
//
//		- if there are no customers, the barber falls asleep in the chair
//		- a customer must wake the barber if he is asleep
//		- if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and
//		  sits in an empty chair if it's available
//		- when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers
//		  and falls asleep if there are none
// 		- shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is
//	      empty
//		- after the shop is closed and there are no clients left in the waiting area, the barber
//		  goes home
//
// The Sleeping Barber was originally proposed in 1965 by computer science pioneer Edsger Dijkstra.
//
// The point of this problem, and its solution, was to make it clear that in a lot of cases, the use of
// semaphores (mutexes) is not needed.

package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

// variables
var seatingCapacity = 5
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 5 * time.Second

func main() {
	// seed out random bumber generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("--------------------------")

	// create channel if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChane := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		ClientsChan:     clientChan,
		NumberOfBarber:  0,
		BarbersDoneChan: doneChane,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	lstBarbers := []string{"Faker", "Guma", "Owner", "Keria", "Doran"}

	// add barbers
	for _, barber := range lstBarbers {
		shop.addBarber(barber)
	}
	//shop.addBarber("Faker")
	//shop.addBarber("Guma")
	//shop.addBarber("Owner")
	//shop.addBarber("Keria")
	//shop.addBarber("Doran")

	ctx, cancel := context.WithTimeout(context.Background(), timeOpen)
	defer cancel()

	// start the barbershop as a goroutine
	//shopClosing := make(chan bool)
	closed := make(chan bool)

	go scheduleShopClosing(ctx, &shop, closed)

	// add clients
	go generateClients(ctx, &shop)

	// block into the barbershop is closed
	<-closed
}

func scheduleShopClosing(ctx context.Context, shop *BarberShop, closed chan bool) {
	<-ctx.Done()
	shop.closeShopForDay()
	closed <- true
}

func generateClients(ctx context.Context, shop *BarberShop) {
	i := 1
	for {
		// get a random number with average arrival rate
		randomMilliseconds := rand.Int() % (2 * arrivalRate)
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(randomMilliseconds) * time.Millisecond):
			shop.addClient(fmt.Sprintf("Client #%d", i))
			i++
		}
	}
}
