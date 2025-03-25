package devhub

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// Client structure representing a GitHub API client
type Client struct {
	HTTPClient *http.Client
}

// NewClient create a GitHub API client
func NewClient() *Client {
	cookieJar, _ := cookiejar.New(nil)
	client := http.Client{
		Timeout:   time.Minute,
		Jar:       cookieJar,
		Transport: http.DefaultTransport.(*http.Transport).Clone(),
	}
	client.Transport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	return &Client{
		HTTPClient: &client,
	}
}

var HandleRequestApiInditex = func(c *Client, url string, method string, password string, body map[string]interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, _ := http.NewRequest(method, url, b)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("itx-apiKey", password)
	req.Header.Add("user-agent", "argo-workflows-inditex")
	resp, err := c.HTTPClient.Do(req)
	if resp == nil {
		resp = new(http.Response)
		// Generate a placeholder body, to avoid nil dereference when
		// running resp.Body.Close() in the caller
		resp.Body = io.NopCloser(bytes.NewBufferString(""))
		return resp, err
	}
	return resp, err
}
