// Let's imagine a restaurant where:

// Waiters (senders) take orders from customers and send them to the kitchen

// Chefs (receivers) prepare the orders as they come in

// Manager (main goroutine) coordinates everything

// Implementation with Channels

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Dish   string
	Table  int
	Status string
}

func main() {
	// Create a buffered channel for orders (can hold up to 5 orders)
	orderChannel := make(chan Order, 5)
	
	// WaitGroup to track workers
	var wg sync.WaitGroup

	// Start 3 chefs (receivers)
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go chef(i, orderChannel, &wg)
	}

	// Start 2 waiters (senders)
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		go waiter(i, orderChannel, &wg)
	}

	// Simulate restaurant operation for some time
	time.Sleep(30 * time.Second)
	
	// Close the kitchen (channel)
	close(orderChannel)
	
	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("Restaurant closed for the day")
}

// Waiter (sender) function
func waiter(id int, orderChan chan<- Order, wg *sync.WaitGroup) {
	defer wg.Done()
	
	dishes := []string{"Pizza", "Pasta", "Salad", "Steak", "Soup"}
	
	for i := 1; ; i++ {
		// Simulate time between taking orders
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
		
		order := Order{
			ID:    i,
			Dish:  dishes[rand.Intn(len(dishes))],
			Table: rand.Intn(10) + 1,
		}
		
		select {
		case orderChan <- order:
			fmt.Printf("Waiter %d: Sent order #%d for %s (Table %d)\n", id, order.ID, order.Dish, order.Table)
		default:
			fmt.Printf("Waiter %d: Kitchen overloaded! Can't send order #%d\n", id, order.ID)
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}
		
		// Stop if kitchen is closed (channel closed)
		if i > 20 { // Safety to prevent infinite loop in example
			break
		}
	}
}

// Chef (receiver) function
func chef(id int, orderChan <-chan Order, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for order := range orderChan {
		fmt.Printf("Chef %d: Started cooking %s for table %d\n", id, order.Dish, order.Table)
		
		// Simulate cooking time
		cookTime := time.Duration(rand.Intn(5)+2) * time.Second
		time.Sleep(cookTime)
		
		fmt.Printf("Chef %d: Finished %s for table %d (took %v)\n", id, order.Dish, order.Table, cookTime)
	}
	
	fmt.Printf("Chef %d: No more orders, going home\n", id)
}