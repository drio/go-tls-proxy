package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func SecureServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Secure Hello World.\n"))
}

func main() {
  port := "443"
	if len(os.Args) == 2 {
    port = os.Args[1]
	}
	http.HandleFunc("/", SecureServer)
	fmt.Printf("Proxying, listening on :%s\n", port)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
