package plutus

import "net/http"

type DepoAccount struct {
	ID     int    `json:"id"`
	Owner  int    `json:"owner_id"`
	Title  string `json:"title"`
	Code   string `json:"code"`
	Broker string `json:"broker"`
}

type BankAccount struct {
	ID           int    `json:"id"`
	Owner        int    `json:"owner_id"`
	Title        string `json:"title"`
	CurrencyRate int    `json:"currency"`
	Code         string `json:"code"`
	Broker       string `json:"broker"`
}

type OwnerAccounts struct {
	Owner int           `json:"owner_title"`
	Depo  []DepoAccount `json:"depo"`
	Bank  []BankAccount `json:"bank"`
}

func NewDepoAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func UpdateDepoAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func GetDepoAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func DeleteDepoAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func NewBankAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func UpdateBankAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func GetBankAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func DeleteBankAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func GetAccountsByOwner(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}
