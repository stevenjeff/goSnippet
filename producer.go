package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(header string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s:  %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string) {
	for {
		message := <-channel
		fmt.Println(message)
	}
}
