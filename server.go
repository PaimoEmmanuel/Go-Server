package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "POST" {
		http.Error(writer, "Method not supported", http.StatusNotFound)
		return
	}
	err := request.ParseForm()
	if err != nil {
		fmt.Println("An error occured: ", err)
		return
	}
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Println("name:", name, "\n", "address:", address)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "Route not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "Method not supported", http.StatusNotFound)
		return
	}
	response := "Hello there!"
	//writer.t
	writer.Write([]byte(response))
}
