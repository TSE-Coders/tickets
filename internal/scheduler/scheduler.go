package scheduler

import (
	"fmt"
	"math/rand"
	"time"
)

type ScheduledFunc func() error

type Schedule struct {
	maxInterval int
	random      bool
	fn          ScheduledFunc
}

func New(maxInterval int, isRandom bool, fn ScheduledFunc) *Schedule {
	return &Schedule{
		maxInterval: maxInterval,
		random:      isRandom,
		fn:          fn,
	}
}

func (s Schedule) Run() {
	go func() {
		randomNumberGenerator := rand.New(rand.NewSource(time.Now().Unix()))
		for {
			wait := s.maxInterval
			if s.random {
				wait = randomNumberGenerator.Intn(int(s.maxInterval))
			}
			time.Sleep(time.Second * time.Duration(wait))
			err := s.fn()
			if err != nil {
				fmt.Printf("failed to run scheduled function: %s\n", err.Error())
			}
		}
	}()
}
