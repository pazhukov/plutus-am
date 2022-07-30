package plutus

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/asset/", NewAsset).Methods("POST")
	router.HandleFunc("/asset/{id}", GetAsset).Methods("GET")
	router.HandleFunc("/assets/{page}", GetAssets).Methods("GET")
	router.HandleFunc("/assets/", GetAssetsStartPage).Methods("GET")
	router.HandleFunc("/code/", AddNewCode).Methods("POST")

	log.Fatal(http.ListenAndServe(":11000", router))

}
