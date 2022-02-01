package main

import (
	"fmt"
	"sync"
)

type hostRequest int

const (
	askPermission hostRequest = iota
	grantPermission
	denyPermission
	finishedEating
)

type Host struct {
	sync.Mutex
	concurrentEaters int
}

func (h Host) watchEaters(askPermissionChannel chan hostRequest, tellOKChannel chan hostRequest, tellFinishedChannel chan hostRequest) {

	for {
		select {
		case <-askPermissionChannel:
			h.Lock()
			if h.concurrentEaters < 2 {
				tellOKChannel <- grantPermission
				h.concurrentEaters++
				fmt.Println("Num concurrently eating: ", h.concurrentEaters)
			} else {
				tellOKChannel <- grantPermission
			}
			h.Unlock()

		case <-tellFinishedChannel:
			h.Lock()
			if h.concurrentEaters > 0 {
				h.concurrentEaters--
			}
			h.Unlock()
		}
	}
}

type ChopStik struct{ sync.Mutex }

type Philosopher struct {
	leftChopstik, rightChopstik *ChopStik
	number, timesEaten          int
}

func (p *Philosopher) eat(askPermissionChannel chan hostRequest, tellOKChannel chan hostRequest, tellFinishedChannel chan hostRequest, wg *sync.WaitGroup) {

	for {
		// block until request been read
		askPermissionChannel <- askPermission

		// block until told
		permissionResult := <-tellOKChannel

		if permissionResult == grantPermission {
			p.leftChopstik.Lock()
			p.rightChopstik.Lock()

			fmt.Println("started eating ", p.number)
			p.timesEaten++
			fmt.Println("finished eating ", p.number)

			p.rightChopstik.Unlock()
			p.leftChopstik.Unlock()

			// tell host finished eating to decrement the ref counter
			tellFinishedChannel <- finishedEating

			if p.timesEaten > 2 {
				fmt.Println("Ive eaten 3 times - I'm full! - done")
				wg.Done()
				break
			}
		}
	}

}

func main() {

	var wg sync.WaitGroup

	askPermissionChannel := make(chan hostRequest)
	tellOKChannel := make(chan hostRequest)
	tellFinishedChannel := make(chan hostRequest)

	chopStiks := make([]*ChopStik, 5)
	for i := 0; i < 5; i++ {
		chopStiks[i] = new(ChopStik)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{chopStiks[i], chopStiks[(i+1)%5], i, 0}
	}

	host := Host{}

	// Rob - the host is outside of the done channel
	go host.watchEaters(askPermissionChannel, tellOKChannel, tellFinishedChannel)

	wg.Add(5) // Rob - the 5 could be consts etc.
	// infinite loop until a 3 times each
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(askPermissionChannel, tellOKChannel, tellFinishedChannel, &wg)
	}
	wg.Wait()
	//https://go.dev/blog/pprof
}
