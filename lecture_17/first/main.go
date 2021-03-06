package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8888", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to go server http from Home page`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to go server http from about`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to go server http from contact`)
}
