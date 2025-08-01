package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of all philosophers
var philosophers = []Philosopher{
	{name: "User 1", leftFork: 4, rightFork: 0},
	{name: "User 2", leftFork: 0, rightFork: 1},
	{name: "User 3", leftFork: 1, rightFork: 2},
	{name: "User 4", leftFork: 2, rightFork: 3},
	{name: "User 5", leftFork: 3, rightFork: 4},
}

// Define a few variables.
var hunger = 3                  // how many times a philosopher eats
var eatTime = 1 * time.Second   // how long it takes to eatTime
var thinkTime = 3 * time.Second // how long a philosopher thinks
var sleepTime = 1 * time.Second // how long to wait when printing things out

var orderMutex sync.Mutex
var orderFinished []string

func main() {
	// print out a welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("===========================")
	fmt.Println("The table is empty")

	time.Sleep(sleepTime)

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")
	fmt.Println("Order: ", strings.Join(orderFinished, ","))
}

func dine() {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal.
	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup,
	forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	//eat three time
	for i := hunger; i > 0; i-- {
		//get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Println(philosopher.name, " is satisfied.")
	fmt.Println(philosopher.name, " left the table.")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
