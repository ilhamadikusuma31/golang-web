package main

import (
	"fmt"
	"golang-web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",handler.HomeHandler)
	mux.HandleFunc("/produk", handler.ProdukHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/proses", handler.Proses)
	mux.HandleFunc("/pembeli",handler.Pembeli)


	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))

	log.Println("server berjalan di port 9000")
	fmt.Println("tes ajah")
	
	err := http.ListenAndServe(":9000",mux)
	log.Fatal(err)
}

