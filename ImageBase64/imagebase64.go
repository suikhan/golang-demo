package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("e:\\a.txt")
	if err != nil {
		fmt.Print(err)
	}
	encodeString := string(b)
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		fmt.Print(err)
	}
	err = ioutil.WriteFile("e:\\a.bmp", decodeBytes, 0644)
	if err != nil {
		fmt.Print(err)
	}
}
