package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var max_productions int

var producers_running bool
var buffer []int

type Consumer struct {
	consuming bool
}

type Producer struct {
	added_items int
}

/*
	The thread that removes from the buffer.
*/
func consumer() {
	consuming := true

	for consuming {
		if len(buffer) > 0 {
			buffer = buffer[:len(buffer)-1]
		}
		fmt.Println("Consumer: ", buffer)
		time.Sleep(time.Duration(rand.Intn(20)))
		if !producers_running && len(buffer) == 0 {
			consuming = false
		}
	}

	fmt.Println("Consumer Stoped!")
	wg.Done()
}

/*
	The thread that adds to the buffer.

	Args:
		maxAdditionOfItems (int): The maximum number of times that the producer will add to the buffer.
*/
func producer(maxAdditionOfItems int) {
	added_items := 0

	for added_items < maxAdditionOfItems {
		buffer = append(buffer, rand.Intn(100))
		fmt.Println("Producer: ", buffer)
		time.Sleep(time.Duration(rand.Intn(10)))
		added_items++
	}

	producers_running = false
	fmt.Println("Producer Stoped!")
	wg.Done()
}

func main() {
	producers_running = true
	max_productions = 100

	numberProducers := 2
	numberConsumers := 3

	// Determinates the number of threads that the algorithm need to wait the execution.
	wg.Add(numberProducers + numberConsumers)

	// Created the producers.
	for i := 0; i < numberProducers; i++ {
		go producer(max_productions)
	}

	// Waits so the buffer has resources.
	time.Sleep(500)

	// Created the consumers.
	for i := 0; i < numberConsumers; i++ {
		go consumer()
	}

	wg.Wait()
}
