package plutus

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DepoAccount struct {
	ID     int    `json:"id"`
	Owner  int    `json:"owner_id"`
	Title  string `json:"title"`
	Code   string `json:"code"`
	Broker string `json:"broker"`
}

type BankAccount struct {
	ID       int    `json:"id"`
	Owner    int    `json:"owner_id"`
	Title    string `json:"title"`
	Currency int    `json:"currency"`
	Code     string `json:"code"`
	Broker   string `json:"broker"`
}

type OwnerAccounts struct {
	Owner int           `json:"owner_title"`
	Depo  []DepoAccount `json:"depo"`
	Bank  []BankAccount `json:"bank"`
}

func NewDepoAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	var input DepoAccount
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Owner == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Owner can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Title == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Code == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Code can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Broker == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Broker can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	ownerID, status := DepoAccountNewItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "New Depo Account with ID = " + strconv.Itoa(ownerID) + " added"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	}

}

func UpdateDepoAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	var input DepoAccount
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.ID == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account ID can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Owner == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Owner can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Title == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Code == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Code can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Broker == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Broker can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	status := DepoAccountUpdateItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = " Depo Account with ID = " + strconv.Itoa(input.ID) + " updated"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 200
		info.Message = " Depo Account with ID = " + strconv.Itoa(input.ID) + " not found"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	}

}

func GetDepoAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	depoAccountID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert Depo Account ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	item, status := DepoAccountGetItem(depoAccountID)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Depo Account not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		json.NewEncoder(w).Encode(item)
	}

}

func DeleteDepoAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	depoAccountID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert Depo Account ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	status := DepoAccountDeleteItem(depoAccountID)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Depo Account not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "Depo Account deleted"
		json.NewEncoder(w).Encode(info)
	}

}

func NewBankAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	var input BankAccount
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Owner == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Owner can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Title == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Currency == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Currency can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Code == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Code can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Broker == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Depo Account Broker can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	ownerID, status := BankAccountNewItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "New Bank Account with ID = " + strconv.Itoa(ownerID) + " added"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	}

}

func UpdateBankAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	var input BankAccount
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.ID == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account ID can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Owner == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Owner can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Title == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Currency == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Currency can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Code == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Code can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Broker == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Bank Account Broker can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	status := BankAccountUpdateItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = " Bank Account with ID = " + strconv.Itoa(input.ID) + " updated"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 200
		info.Message = " Bank Account with ID = " + strconv.Itoa(input.ID) + " not found"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	}

}

func GetBankAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	bankAccountID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert Bank Account ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	item, status := BankAccountGetItem(bankAccountID)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Bank Account not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		json.NewEncoder(w).Encode(item)
	}

}

func DeleteBankAccount(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	bankAccountID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert Bank Account ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	status := BankAccountDeleteItem(bankAccountID)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Bank Account not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "Bank Account deleted"
		json.NewEncoder(w).Encode(info)
	}

}

func GetAccountsByOwner(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	ownerID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert owner ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	item, status := OwnerGetAccounts(ownerID)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Owner not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		json.NewEncoder(w).Encode(item)
	}

}
