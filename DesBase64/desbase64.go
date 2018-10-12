//注意base64对中文不同编码加解密结果不一样，即UTF-8跟GBK
package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"

	"github.com/axgle/mahonia"
)

//DES加密
func DesEncrypt(origData, key []byte) (string, error) {
	//UTF-8 to GBK
	var enc mahonia.Encoder
	enc = mahonia.NewEncoder("gbk")
	origDataStr := enc.ConvertString(string(origData))
	origData = []byte(origDataStr)

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	//crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	//base64加密
	encodeString := base64.StdEncoding.EncodeToString(crypted)
	return encodeString, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//DES解密
func DesDecrypt(encodeString string, key []byte) (string, error) {
	var dec mahonia.Decoder
	//base64解密
	crypted, err := base64.StdEncoding.DecodeString(encodeString)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)

	origData = ZeroUnPadding(origData)
	//GBK to UTF-8
	dec = mahonia.NewDecoder("gbk")
	origDataStr := dec.ConvertString(string(origData))

	return origDataStr, nil
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func main() {
	fmt.Println("DES与BASE64相结合的加密方式")
	fmt.Println("==========================")
	c := []byte("@suikhan")
	var str string = string(c[:])
	fmt.Println("请输入密钥(", str, ")")
	fmt.Scanln(&str)
	if str != "" {
		c = []byte(str)
	}
	fmt.Println("请输入需要加密的字符串：")
	str = ""
	fmt.Scanln(&str)
	b := []byte("杨雨田(suikhan@126.com)")
	if str != "" {
		b = []byte(str)
	}
	fmt.Println("==========================")
	a, _ := DesEncrypt(b, c)
	fmt.Println("密文：", a)

	q, _ := DesDecrypt(a, c)
	fmt.Println("明文：", q)
}
