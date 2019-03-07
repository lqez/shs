package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := "5000"

    if len(os.Args) > 1 {
        port = os.Args[1]
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("-", r.URL.Path)
        http.ServeFile(w, r, r.URL.Path[1:])
    })

    log.Println(fmt.Sprintf("Serving HTTP on http://localhost:%s", port))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
