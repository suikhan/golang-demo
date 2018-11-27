package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func main() {
	encodeString := "123456"
	h := md5.New()
	h.Write([]byte(encodeString)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)

	decodeBytes := base64.StdEncoding.EncodeToString(cipherStr)

	fmt.Println(decodeBytes)

}
