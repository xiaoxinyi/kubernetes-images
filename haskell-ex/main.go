package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("NAME_TO_GREET")
	fmt.Printf("hello %s\n", name)

}