package plutus

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/asset/", NewAsset).Methods("POST")
	router.HandleFunc("/asset/", UpdateAsset).Methods("PUT")
	router.HandleFunc("/asset/{id}", GetAsset).Methods("GET")
	router.HandleFunc("/asset/{id}", DeleteAsset).Methods("DELETE")
	router.HandleFunc("/assets/{page}", GetAssets).Methods("GET")
	router.HandleFunc("/assets/", GetAssetsStartPage).Methods("GET")
	router.HandleFunc("/code/", AddNewCode).Methods("POST")
	router.HandleFunc("/code/", DeleteCode).Methods("DELETE")
	router.HandleFunc("/currency/", NewCurrency).Methods("POST")
	router.HandleFunc("/currency/", UpdateCurrency).Methods("PUT")
	router.HandleFunc("/currency/{id}", GetCurrency).Methods("GET")
	router.HandleFunc("/currency/{id}", DeleteCurrency).Methods("DELETE")
	//router.HandleFunc("/currency/rates", GetLastCurrencyRate).Methods("GET")
	//router.HandleFunc("/currency/rates/{date}", GetCurrencyRateByDate).Methods("GET")

	router.HandleFunc("/load/currency-rates", LoadCurrencyRates).Methods("GET")

	log.Fatal(http.ListenAndServe(":11000", router))

}
