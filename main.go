package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",homeHandler)
	mux.HandleFunc("/produk", produkHandler)

	log.Println("server berjalan di port 9000")
	
	err := http.ListenAndServe(":9000",mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}

	w.Write([]byte("<h1>tes aja</h1>"))
}

func produkHandler(w http.ResponseWriter, r *http.Request){

	tangkap_id := r.URL.Query().Get("id")
	id,err := strconv.Atoi(tangkap_id)

	if err != nil || id < 1{
		http.NotFound(w, r)
	}

	fmt.Fprintf(w,"Produk dengan id %d",id)
}