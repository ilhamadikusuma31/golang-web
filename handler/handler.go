package handler

import (
	"golang-web/entity"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	templ, err := template.ParseFiles("views/home.html", "views/layout.html")
	if err != nil {
		log.Println("error for programmer: ", err)
		http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
		return
	}

	//data
	data := map[string]string{
		"title":   "ini judul konten di home",
		"content": "ini konten di halaman home",
	}

	err = templ.Execute(w, data)
	if err != nil {
		log.Println("error for programmer: ", err)
		http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
		return
	}
	// w.Write([]byte("<h1>tes aja</h1>"))
}

func ProdukHandler(w http.ResponseWriter, r *http.Request) {

	tangkap_id := r.URL.Query().Get("id")
	id, err := strconv.Atoi(tangkap_id)

	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	// fmt.Fprintf(w, "Produk dengan id %d", id)

	templ, err := template.ParseFiles("views/produk.html", "views/layout.html")
	if err != nil {
		log.Println("error for programmer: ", err)
		http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
		return
	}
	// data := map[string]interface{}{
	// 	"title" : "ini judul konten di home",
	// 	"content": id,
	// }

	//data := entity.Produk{ID: 1, Nama: "iphone x", Harga: 12000000}
	data := []entity.Produk{
		{ID: 1, Nama: "iphone x", Harga: 12000000},
		{ID: 2, Nama: "samsung galaxy j12", Harga: 1200000},
		{ID: 3, Nama: "esia hidayah", Harga: 15000},
		{ID: 3, Nama: "nokia 3310", Harga: 200000}}

	err = templ.Execute(w, data)
	if err != nil {
		log.Println("error for programmer: ", err)
		http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
		return
	}
}
