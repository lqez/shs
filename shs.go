package main

import (
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("-", r.URL.Path)
        http.ServeFile(w, r, r.URL.Path[1:])
    })

    log.Println("Serving HTTP on http://localhost:5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}
