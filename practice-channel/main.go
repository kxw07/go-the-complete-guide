package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	stringChannel := make(chan string)
	errorChannel := make(chan error)

	go returnString("Hello, World!", stringChannel)
	go returnError(errors.New("an Error"), errorChannel)

	for range 2 {
		select {
		case str := <-stringChannel:
			fmt.Println(str)
		case err := <-errorChannel:
			fmt.Println(err)
		}
	}
}

func returnString(str string, channel chan string) {
	time.Sleep(3 * time.Second)
	channel <- str
}

func returnError(err error, channel chan error) {
	time.Sleep(1 * time.Second)
	channel <- err
}
