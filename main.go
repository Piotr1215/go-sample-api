package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

func home(res http.ResponseWriter, r *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(res, "HOME Page")
}

func about(res http.ResponseWriter, r *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(res, "ABOUT Page")
}

func version(res http.ResponseWriter, r *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(res, "VERSION Page")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/version", version)
	fmt.Println("go to: http://localhost:8080/")
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
