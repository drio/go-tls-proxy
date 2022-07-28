package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, this is the service.\n"))
}

func main() {
  port := "8080"
	if len(os.Args) == 2 {
    port = os.Args[1]
	}
	http.HandleFunc("/", handleRequest)
	fmt.Printf("service running on: %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
