package src

import "sync"

// Environment simulates a computational en environment.
type Environment struct {
	Wg               sync.WaitGroup
	MaxProductions   int
	ProducersRunning bool
	Buffer           []int
}
