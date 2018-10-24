package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	SayHello("Nihao")
}

func SayHello(name string) string {
	return "Hello," + name
}
