package config

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"

	log "github.com/sirupsen/logrus"
)

func Compress(str string) string {
	if str == "" {
		return ""
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		log.Errorf("failed write Gzip on Compress: %v", err)
		return ""
	}
	if err := gz.Close(); err != nil {
		log.Errorf("failed to close Gzip on Compress: %v", err)
		return ""
	}
	strEncoded := base64.StdEncoding.EncodeToString(b.Bytes())
	return strEncoded
}

func Decompress(str string) string {
	if str == "" {
		return ""
	}
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Errorf("failed to decode base64 on Decompress: %v", err)
		return ""
	}
	rdata := bytes.NewReader(data)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		log.Errorf("failed to create Gzip reader on Decompress: %v", err)
		return ""
	}
	s, err := io.ReadAll(r)
	if err != nil {
		log.Errorf("failed to read Gzip on Decompress: %v", err)
		return ""
	}
	return string(s)
}
