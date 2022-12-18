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
	router.HandleFunc("/currency-rates", GetLastCurrencyRate).Methods("GET")
	router.HandleFunc("/currency-rates/{date}", GetCurrencyRateByDate).Methods("GET")
	router.HandleFunc("/load/currency-rates", LoadCurrencyRates).Methods("GET")
	router.HandleFunc("/load/currency-rates/{date}", LoadCurrencyRatesByDate).Methods("GET")
	router.HandleFunc("/owner/", NewOwner).Methods("POST")
	router.HandleFunc("/owner/", UpdateOwner).Methods("PUT")
	router.HandleFunc("/owner/{id}", GetOwner).Methods("GET")
	router.HandleFunc("/owner/{id}", DeleteOwner).Methods("DELETE")
	router.HandleFunc("/owners/", GetOwnerList).Methods("GET")
	router.HandleFunc("/account/depo/", NewDepoAccount).Methods("POST")
	router.HandleFunc("/account/depo/", UpdateDepoAccount).Methods("PUT")
	router.HandleFunc("/account/depo/{id}", GetDepoAccount).Methods("GET")
	router.HandleFunc("/account/depo/{id}", DeleteDepoAccount).Methods("DELETE")
	router.HandleFunc("/account/bank/", NewBankAccount).Methods("POST")
	router.HandleFunc("/account/bank/", UpdateBankAccount).Methods("GET")
	router.HandleFunc("/account/bank/{id}", GetBankAccount).Methods("GET")
	router.HandleFunc("/account/bank/{id}", DeleteBankAccount).Methods("DELETE")
	router.HandleFunc("/account/owner/{id}", GetAccountsByOwner).Methods("GET")
	router.HandleFunc("/operation/inout/depo/", NewInOutDepo).Methods("POST")
	router.HandleFunc("/operation/inout/depo/", UpdateInOutDepo).Methods("PUT")
	router.HandleFunc("/operation/inout/depo/{id}", GetInOutDepo).Methods("GET")
	router.HandleFunc("/operation/inout/depo/{id}", DeleteInOutDepo).Methods("DELETE")
	router.HandleFunc("/operation/inout/bank/", NewInOutBank).Methods("POST")
	router.HandleFunc("/operation/inout/bank/", UpdateInOutBank).Methods("PUT")
	router.HandleFunc("/operation/inout/bank/{id}", GetInOutBank).Methods("GET")
	router.HandleFunc("/operation/inout/bank/{id}", DeleteInOutBank).Methods("DELETE")
	// router.HandleFunc("/report/inout/depo/{date1}/{date2}", ReportInOutDepo).Methods("GET")
	// router.HandleFunc("/report/inout/depo/{date1}/{date2}", ReportInOutDepo).Methods("GET")

	log.Fatal(http.ListenAndServe(":11000", router))

}
