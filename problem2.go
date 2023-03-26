package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenChannel := make(chan int)
	oddChannel := make(chan int)

	go getEvenData(numbers, evenChannel)
	go getOddData(numbers, oddChannel)

	evenNumbers := <-evenChannel
	oddNumbers := <-oddChannel

	fmt.Println("Even numbers:", evenNumbers)
	fmt.Println("Odd numbers:", oddNumbers)
}

func getEvenData(numbers []int, ch chan<- int) {
	var evenNumbers []int

	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Println("Even numbers are :", evenNumbers)

	ch <- len(evenNumbers)
}

func getOddData(numbers []int, ch chan<- int) {
	var oddNumbers []int

	for _, num := range numbers {
		if num%2 != 0 {
			oddNumbers = append(oddNumbers, num)
		}
	}
	fmt.Println("Odd numbers are :", oddNumbers)

	ch <- len(oddNumbers)
}
