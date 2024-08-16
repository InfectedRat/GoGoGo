package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://127.0.0.1:5555/get")
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("error", err)
		return
	}

	fmt.Printf("%s", data)
}
