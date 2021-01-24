package main

import "fmt"

type fifo struct{}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by First In First Out strategy.")
}
