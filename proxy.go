package main

import (
  "flag"
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
  "os"
)

const version = "0.0.6"

var (
  flgHelp     bool
  flgProxyUrl string
  flgPort string
  flgCert string
  flgKeys string
)

func parseCmdLineFlags() {
  flag.BoolVar(&flgHelp, "help", false, "if true, show help")
  flag.StringVar(&flgProxyUrl, "proxy-url", "", "Url to proxy to. E.g. http://localhost:8080")
  flag.StringVar(&flgPort, "port", "443", "Port for the proxy to listen to (default: 443)")
  flag.StringVar(&flgCert, "cert", "./server.crt", "path to certificate (default: ./server.crt)")
  flag.StringVar(&flgKeys, "keys", "./server.key", "path to keys (default: ./server.key)")
  flag.Parse()
}

func main() {
  parseCmdLineFlags()
  if flgHelp {
    flag.Usage()
    os.Exit(0)
  }

  if flgProxyUrl == "" {
    flag.Usage()
    os.Exit(0)
  }

  remote, err := url.Parse(flgProxyUrl)
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

  fmt.Printf("Proxying, listening on port:%s proxy to url:%s\n", flgPort, flgProxyUrl)
  err = http.ListenAndServeTLS(fmt.Sprintf(":%s", flgPort), flgCert, flgKeys, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
