package main

import (
	"fmt",
	"log",
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", httpStatusNotFound)
	}
}


func main() {
	fileServer := http.FileServer(http.Dir('./static'))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf('Server starting at port 8080\n')
	if err := http.ListenAndServer(":8080", nil); err != nil {
		log.Fatal(err)
	}
}


