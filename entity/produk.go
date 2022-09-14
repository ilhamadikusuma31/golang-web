package entity

type Produk struct {
	ID    int
	Nama  string
	Harga int
}

func (p Produk) CekMahalGa() string {
	var status string
	if p.Harga > 1000000 {
		status = "mahal kali bos"
	} else {
		status = "murah bang"
	}
	return status

}