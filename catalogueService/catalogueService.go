package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var catalogue []Sock 

var dev bool
var port string

func main() {

	flag.BoolVar(&dev, "dev", false, "Run in development mode")
	flag.StringVar(&port, "port", "8081", "Port on which to run")
	flag.Parse()

	var file string
	if dev {
		file = "./socks.json"
	} else {
		file = "/config/socks.json"
	}
	loadCatalogue(file)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/catalogue", catalogueHandler)
	router.HandleFunc("/catalogue/{catId}", itemHandler)
	fmt.Printf("Catalogue service running on port %s\n", port)
	http.ListenAndServe(":" + port, router)
}

func catalogueHandler(w http.ResponseWriter, r *http.Request) {
	var data []byte
	var err error

	data, err = json.Marshal(catalogue)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    catId := vars["catId"]

	for _, sock := range catalogue {
		if sock.Id == catId {
			data, _ := json.Marshal(sock)
			w.Write(data)
			return
		}
	}
	w.WriteHeader(404)
}

func loadCatalogue(file string) {
	f, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err)
    }

    json.Unmarshal(f, &catalogue)
    fmt.Printf("Loaded %d items into catalogue.\n", len(catalogue))
}

type Sock struct {
	Id string `json:id`
	Price int `json:price`
	Count int `json:"count"`
	// Size, image, name?
}