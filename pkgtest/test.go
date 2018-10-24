package main

import (
	"fmt"

	"./pkg01"
)

func main() {
	fmt.Println("From main!!!")
	pkg01.PrintString("from package!!!")
}
