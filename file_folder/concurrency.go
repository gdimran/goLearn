package main

import (
	"fmt"
	"sync"
	"time"
)

//run multipule function at a time not parallel but concurrently

func main() {

	//waitgroup will wait for waiting concurrency
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	go cooking("rice", &wg)
	go cooking("curry", &wg)

	wg.Wait()
}

//sending wait group function as refencing
func cooking(msg string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
