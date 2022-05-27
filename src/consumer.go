package src

import (
	"fmt"
	"math/rand"
	"time"
)

type Consumer struct {
	Id        int
	Consuming bool
}

// startConsuming is the thread that removes from the buffer.
func (c *Consumer) StartConsuming(env *Environment) {
	c.Consuming = true

	for c.Consuming {
		if len(env.Buffer) > 0 {
			env.Buffer = env.Buffer[:len(env.Buffer)-1]
		}

		fmt.Printf("Consumer %v: %v\n", c.Id, env.Buffer)

		// Waits a random amount of time to simulate a consuming time.
		time.Sleep(time.Duration(rand.Intn(20)))

		if !env.ProducersRunning && len(env.Buffer) == 0 {
			c.Consuming = false
		}
	}

	fmt.Printf("Consumer %v Stopped!\n", c.Id)
	env.Wg.Done()
}
