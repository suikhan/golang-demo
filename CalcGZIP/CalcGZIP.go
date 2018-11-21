package CalcGZIP

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

func Zip(rawData []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(rawData)
	w.Flush()
	str := b.String()
	return []byte(str)
}

func UnZip(zippedData []byte) []byte {
	var b bytes.Buffer
	b.Write(zippedData)
	r, _ := gzip.NewReader(&b)
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)
	return undatas
}

func EncodeJSON(rawString string) string {
	bytes := []byte(rawString)
	zippedData := Zip(bytes)
	strBase64 := base64.URLEncoding.EncodeToString(zippedData)
	return strBase64
}

func DecodeJSON(zippedString string) string {
	str, _ := base64.URLEncoding.DecodeString(zippedString)
	bytes := []byte(str)
	unzippedData := UnZip(bytes)
	return string(unzippedData)
}
