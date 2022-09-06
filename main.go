package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",helloHandler)
	
	err := http.ListenAndServe(":9000",mux)
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>tes aja</h1>"))
}