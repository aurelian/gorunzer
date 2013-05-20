package main

import (
  "fmt"
  "html"
  "net/http"
)

func main() {

  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  http.Handle("/", http.FileServer(http.Dir("public")))

  http.ListenAndServe(":9090", nil)
}
