package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Path: ", r.URL.Path)
	fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[1:])
	fmt.Println(r.URL.Path)
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
