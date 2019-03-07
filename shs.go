package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := "5000"
    notFoundPage := "404.html"

    if len(os.Args) > 1 {
        port = os.Args[1]
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        filepath := r.URL.Path
        if _, err := os.Stat(filepath[1:]); err == nil {
            log.Println("-", filepath)
            http.ServeFile(w, r, filepath[1:])
        } else if os.IsNotExist(err) {
            log.Println("- Not found -", filepath)
            http.ServeFile(w, r, notFoundPage)
        }
    })

    log.Println(fmt.Sprintf("Serving HTTP on http://localhost:%s", port))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
