package src

import (
	"fmt"
	"math/rand"
	"time"
)

type Producer struct {
	Id         int
	AddedItems int
}

// startProducing is the thread that adds to the buffer.
func (p *Producer) StartProducing(env *Environment) {
	p.AddedItems = 0

	for p.AddedItems < env.MaxProductions {
		env.Buffer = append(env.Buffer, rand.Intn(100))

		fmt.Printf("Producer %v: %v\n", p.Id, env.Buffer)

		// Waits a random amount of time to simulate a producing time.
		time.Sleep(time.Duration(rand.Intn(10)))

		p.AddedItems++
	}

	env.ProducersRunning = false

	fmt.Printf("Producer %v Stopped!\n", p.Id)
	env.Wg.Done()
}
