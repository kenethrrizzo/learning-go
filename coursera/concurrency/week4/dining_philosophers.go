package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	number                        int
	leftChopStick, rightChopStick *ChopStick
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		p.leftChopStick.Lock()
		p.rightChopStick.Lock()

		fmt.Printf("starting to eat -> %d\n", p.number)
		time.Sleep(time.Second)

		p.leftChopStick.Unlock()
		p.rightChopStick.Unlock()

		fmt.Printf("finishing eating -> %d\n", p.number)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, chopSticks[i], chopSticks[(i+1)%5]}
		wg.Add(1)
		go philosophers[i].eat(&wg)
	}

	wg.Wait()
}
