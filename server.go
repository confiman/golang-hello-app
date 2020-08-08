package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello Hepsiburada from @Yavuz")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":11130", nil)
}

