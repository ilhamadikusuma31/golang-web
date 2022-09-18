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

	err = templ.Execute(w, data) //eksekusi sambil bawa data
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

	templ, err := template.ParseFiles("views/produk.html", "views/layout.html")
	if err != nil {
		log.Println("error for programmer: ", err)
		http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
		return
	}
	data := []entity.Produk{
		{ID: 1, Nama: "iphone x", Harga: 12000000},
		{ID: 2, Nama: "samsung galaxy j12", Harga: 1200000},
		{ID: 3, Nama: "esia hidayah", Harga: 15000},
		{ID: 4, Nama: "nokia 3310", Harga: 200000}}

	err = templ.Execute(w, data)
	if err != nil {
		log.Println("error for programmer: ", err)
		http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah method GET"))
	case "POST":
		w.Write([]byte("ini adalah method POST"))
	default:
		http.Error(w, "error boskuh gada method yg sesuai", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("views/form.html", "views/layout.html")
		if err != nil {
			log.Println("error for programme: ", err)
			http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
			return
		}
		err = templ.Execute(w, nil)
		if err != nil {
			log.Println("error for programmer: ", err)
			http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
			return
		}

		return
	}
	http.Error(w, "error bos", http.StatusInternalServerError)
}

func Proses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("error for programmer: ", err)
			http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
			return
		}
		// nama := r.Form.Get("nama")
		//data := entity.GetData()
		//data := []entity.Penampung{
		//
		//
		//entity.WriteData(data)
		return
	}

}

func Pembeli(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pembeli := entity.GetData()
		templ, err := template.ParseFiles("views/pembeli.html", "views/layout.html")
		if err != nil {
			log.Println("error for programme: ", err)
			http.Error(w, "Maaf halaman tidak ditemukan", http.StatusInternalServerError)
			return
		}
		log.Println(pembeli)
		templ.Execute(w, pembeli.Data)
	}
}
