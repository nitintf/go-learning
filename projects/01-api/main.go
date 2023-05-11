package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello world")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	fmt.Fprintf(w, "POST request succesfull\n")

	name := r.FormValue("name")
	age := r.FormValue("age")

	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Age: %v", age)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Started GO server on 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
