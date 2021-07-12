package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type data struct {
	k string
	v int
}

//var random = rand.Int()

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		//choice := rng.Intn(len(l))
		//fmt.Printf("choice: %d\n", choice)
		b[i] = letterRunes[rand.Int() % len(letterRunes)]
		//b[i] = l[choice]
	}
	return string(b)
}

func listener(input <-chan []data) error {
	for {
		all := []data{}
		select {
		case i := <- input:
			for _, item := range i {
				all = append(all, item)
			}
			fmt.Printf("Received %d items, full size: %d\n", len(i), len(all))
		}
	}
	return nil
}

func main() {
	// for randomness
	m1 := rand.NewSource(time.Now().UnixNano())
	mRng := rand.New(m1)
	rngs := []*rand.Rand{}
	MAX_WORKERS := 4

	// need to make a few rngs because its not safe for concurrent use
	for i := 0; i < MAX_WORKERS + 1; i++ {
		s1 := rand.NewSource(time.Now().UnixNano()+int64(i+1000))
		rngs = append(rngs, rand.New(s1))
	}

	// send to the other queue
	dataQueue := make(chan []data)

	// listen for data
	go listener(dataQueue)

	// wait group for each iteration
	var wg sync.WaitGroup
	for {
		// create a random number of workers between 1-4
		numWorkers := mRng.Intn(4) + 1 // must be at least 1
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func(q chan []data, wg *sync.WaitGroup) {
				//randstr := RandStringRunes(rand.New(m1), 8, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
				defer wg.Done()
				// random number of items
				numItems := rngs[i].Intn(100) + 1
				items := []data{}
				for j := 0; j < numItems; j++ {
					// make random key of length 8
					items = append(items, data{
						k: RandStringRunes(8),
						v: rand.New(m1).Intn(10000),
					})
				}
				q <- items
			}(dataQueue, &wg)
		}
		wg.Wait()
		//time.Sleep(time.Duration(time.Millisecond * 500))
	}
}
