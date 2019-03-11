package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := "5000"
    indexPage := "index.html"
    notFoundPage := "404.html"

    if len(os.Args) > 1 {
        port = os.Args[1]
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        filepath := r.URL.Path[1:]

        if len(filepath) == 0 {
            filepath = indexPage
        }

        if _, err := os.Stat(filepath); err == nil {
            log.Println("-", filepath)
            http.ServeFile(w, r, filepath)
        } else if os.IsNotExist(err) {
            log.Println("- Not found -", filepath)
            http.ServeFile(w, r, notFoundPage)
        }
    })

    log.Println(fmt.Sprintf("Serving HTTP on http://localhost:%s", port))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
