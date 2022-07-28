package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	port := "443"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	remote, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.Host = remote.Host
			w.Header().Set("X-Stuff", "Foo")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	http.HandleFunc("/cert_metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Metrics here"))
	})

	fmt.Printf("Server, listening on :%s\n", port)
	err = http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
