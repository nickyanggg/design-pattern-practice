package main

import "sync"

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go getInstance(wg)
	}
	wg.Wait()
}
