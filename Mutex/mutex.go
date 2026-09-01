package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views += 1

}

func main() {
	mypost := post{views: 0}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mypost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(mypost.views)
}
