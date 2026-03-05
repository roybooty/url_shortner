package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func createShort(w http.ResponseWriter, r *http.Request){
	fmt.Println("i have been accessed by %s", "safwan")
}

func retreiveShort(w http.ResponseWriter, r *http.Request){
	fmt.Println("i have been accessed bro %s", "mee")
}

func deleteShort(w http.ResponseWriter, r *http.Request){
	fmt.Println("i have been accessed bro %s", "mee")
}

func shortStats(w http.ResponseWriter, r *http.Request){
	fmt.Println("okay now %s", "ha")
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(os.Stdout,"i am live %s", "safwan");
}

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/", hello)
	r.HandleFunc("/new", createShort).Methods("POST")
	r.HandleFunc("/retrieve/{id}", retreiveShort).Methods("GET")
	r.HandleFunc("/delete/{id}", deleteShort).Methods("DEL")
	r.HandleFunc("/statistics/{id}", shortStats).Methods("GET")

	http.ListenAndServe(":5500", r)
	fmt.Println("api running successfully")
}
