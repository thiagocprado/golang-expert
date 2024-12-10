package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	// podemos definir além de timeouts outros headers
	c := http.Client{Timeout: time.Microsecond}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
