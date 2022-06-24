package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"io"
	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name{PARAM}", handelName).Methods("GET")
	router.HandleFunc("/bad", handelBad).Methods("GET")
	router.HandleFunc("/data", handelData).Methods("POST")
	router.HandleFunc("/header",  handelHeader).Methods("GET")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
func handelName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, err :=fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
	if err != nil {
		log.Fatalln(err)

	}
}

func handelBad(w http.ResponseWriter, r *http.Request) {
		 w.WriteHeader(http.StatusBadGateway)
	}

func handelData(w http.ResponseWriter, r *http.Request) {
	 d, err := io.ReadAll(r.Body)
	if err == nil {
		fmt.Fprintf(w, "I got message:\nPARAM")
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(d)
}

func handelHeader(w http.ResponseWriter, r *http.Request) {

}
