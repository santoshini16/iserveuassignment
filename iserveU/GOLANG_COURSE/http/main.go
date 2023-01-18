package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logwriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logwriter{}

	io.Copy(lw, resp.Body)
}
func (logwriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just rote this many bytes:", len(bs))
	return len(bs), nil
}
