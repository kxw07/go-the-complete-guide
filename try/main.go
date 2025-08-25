package main

import (
	"fmt"

	"github.com/samborkent/uuidv7"
)

func main() {
	fmt.Println(uuidv7.New().String())
}
