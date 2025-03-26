package config

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"log"
)

func Compress(str string) string {
	if str == "" {
		return ""
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		log.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}
	strEncoded := base64.StdEncoding.EncodeToString(b.Bytes())
	return strEncoded
}

func Decompress(str string) string {
	if str == "" {
		return ""
	}
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := io.ReadAll(r)
	return string(s)
}
