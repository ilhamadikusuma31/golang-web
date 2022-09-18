package entity

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Pembeli struct {
	Name_Entity string
	Data        []*Data
}

type Data struct {
	Nama string
	Umur string
}

func errorHandler(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetData() Pembeli {
	hasil, err := ioutil.ReadFile("./database/pembeli.json")
	errorHandler(err)
	var p Pembeli
	jErr := json.Unmarshal(hasil, &p)
	errorHandler(jErr)
	return p
}

func WriteData(data Pembeli) {
	hasil, err := json.Marshal(data)
	errorHandler(err)
	jErr := ioutil.WriteFile("./database/pembeli.json", hasil, 0775)
	errorHandler(jErr)

}
