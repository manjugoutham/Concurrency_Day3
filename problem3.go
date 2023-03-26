package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomChannel := make(chan int, 3)

	go sendRandomIntegers(randomChannel)

	for i := 0; i < 15; i++ {
		time.Sleep(time.Second)
		select {
		case randomIntegers := <-randomChannel:
			fmt.Println("Received random integer:", randomIntegers)
		default:
			fmt.Println("No data received")
		}
	}
}

func sendRandomIntegers(ch chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		randomIntegersValue := rand.Intn(100)
		ch <- randomIntegersValue
		time.Sleep(time.Second)
	}
}
