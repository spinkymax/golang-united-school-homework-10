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
	router.HandleFunc("/name{PARAM}", handleName).Methods("GET")
	router.HandleFunc("/bad", handleBad).Methods("GET")
	router.HandleFunc("/data", handleData).Methods("POST")
	router.HandleFunc("/headers",  handleHeader).Methods("POST")

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
func handleName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(" Hello, " +vars ["PARAM"]+ "!"))
	
}

func handleBad(w http.ResponseWriter, r *http.Request) {
		 w.WriteHeader(http.StatusInternalServerError)
	}

func handleData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	d, err:= io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
		fmt.Fprintf(w, "I got message:\n"+string(d))
	}

func handleHeader(w http.ResponseWriter, r *http.Request) {
        a := r.Header.Get("a")
	b := r.Header.Get("b")

	aa, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bb, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c := strconv.Itoa(aa + bb)
	w.Header().Add("a+b", fmt.Sprint(c))
	w.WriteHeader(http.StatusOK)
}
