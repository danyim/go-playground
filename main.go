package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var PORT = 8081

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}

func main() {
	fmt.Printf("Server active at http://localhost:%d\r\n", PORT)
	//handleRequests()
	serve()
}
