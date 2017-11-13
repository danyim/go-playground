package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, "Counter: %d", counter)
	mutex.Unlock()
}

func serve() {
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/inc", incrementCounter)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
