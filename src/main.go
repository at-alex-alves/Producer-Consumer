package main

import (
	"time"

	"github.com/at-alex-alves/Producer-Consumer/src/models"
)

func main() {
	environment := src.Environment{
		MaxProductions:   100,
		ProducersRunning: true,
		Buffer:           []int{},
	}

	numberProducers := 2
	numberConsumers := 3

	// Determinates the number of threads that the algorithm need to wait the execution.
	environment.Wg.Add(numberProducers + numberConsumers)

	// Created the producers.
	for currentProducerId := 0; currentProducerId < numberProducers; currentProducerId++ {
		producer := src.Producer{
			Id:         currentProducerId,
			AddedItems: 0,
		}

		go producer.StartProducing(&environment)
	}

	// Waits so the buffer has resources.
	time.Sleep(500)

	// Created the consumers.
	for currentConsumerId := 0; currentConsumerId < numberConsumers; currentConsumerId++ {
		consumer := src.Consumer{
			Id:        currentConsumerId,
			Consuming: false,
		}

		go consumer.StartConsuming(&environment)
	}

	environment.Wg.Wait()
}
