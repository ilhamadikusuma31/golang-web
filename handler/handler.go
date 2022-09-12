package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("<h1>tes aja</h1>"))
}

func ProdukHandler(w http.ResponseWriter, r *http.Request) {

	tangkap_id := r.URL.Query().Get("id")
	id, err := strconv.Atoi(tangkap_id)

	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "Produk dengan id %d", id)
}