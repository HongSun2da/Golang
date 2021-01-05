package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(2 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)

	go push(c)

	timeerChan := time.After(10 * time.Second)
	tickChan := time.Tick(2 * time.Second)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timeerChan:
			fmt.Println("timeout")
			return
		case <-tickChan:
			fmt.Println("tickChan")
		default:
			fmt.Println("Idle")
			time.Sleep(1 * time.Second)
		}
	}
}
