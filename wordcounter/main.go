package main

import (
	"felix/golangprojects/counter"
	"fmt"
	"os"
)

func main() {
	fmt.Println(counter.Count(os.Stdin))
}
