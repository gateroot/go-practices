package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player("A", table)
	go player("B", table)

	table <- Ball
	time.Sleep(1 * time.Second)
	fmt.Println("result: ", <-table)
}

func player(name string, table chan int) {
	for {
		ball := <-table
		ball++
		fmt.Printf("[%s] %d\n", name, ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
