package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://localhost:18888")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
}
