package main

import "sync"

var hits struct {
	sync.Mutex
	n int
}

func main() {
	hits.Lock()
	hits.n++
	hits.Unlock()
}
