package filehandling

import (
	"encoding/json"
	"io/ioutil"
	"log"
)


type Penampung struct{
	Name_Entity string
	Data []*Pembeli
}

type Pembeli struct{
	Nama string
	Umur string
}

func errorHandler(e error)  {
	if e!=nil {
		log.Fatal(e)
	}
}


func GetData() Penampung{
	hasil, err := ioutil.ReadFile("./database/pembeli.json")
	errorHandler(err)
	var p Penampung
	jErr:=json.Unmarshal(hasil, &p)
	errorHandler(jErr)
	return p	
}