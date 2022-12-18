package plutus

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type DepoInOut struct {
	ID        int     `json:"id"`
	Date      string  `json:"inout_date"`
	Operation int     `json:"operation_id"`
	Asset     int     `json:"asset_id"`
	Account   int     `json:"depo_account_id"`
	Count     float64 `json:"count"`
}

type BankInOut struct {
	ID        int     `json:"id"`
	Date      string  `json:"inout_date"`
	Operation int     `json:"operation_id"`
	Account   int     `json:"bank_account_id"`
	Amount    float64 `json:"amount"`
	TypeInOut string  `json:"inout_type"`
}

type OperationDepoInOut struct {
	ID       int         `json:"id"`
	Date     string      `json:"operation_date"`
	Number   string      `json:"operation_number"`
	Type     string      `json:"type"`
	Trade    int         `json:"trade_id"`
	DepoData []DepoInOut `json:"depo"`
	Comment  string      `json:"comment"`
}

type OperationBankInOut struct {
	ID       int         `json:"id"`
	Date     string      `json:"operation_date"`
	Number   string      `json:"operation_number"`
	Type     string      `json:"type"`
	Trade    int         `json:"trade_id"`
	BankData []BankInOut `json:"bank"`
	Comment  string      `json:"comment"`
}

func NewInOutDepo(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	var input OperationDepoInOut
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	_, err_date := time.Parse("2006-01-02", input.Date)
	if err_date != nil {
		var info InfoMessage
		info.Code = 301
		info.Message = "Wrong Operation Date, use format YYYY-MM-DD"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Number == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Operation Mumber can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Type == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Operation Type can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Type != "inout" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Operation Type must be 'inout'"
		json.NewEncoder(w).Encode(info)
		return
	}

	if len(input.DepoData) == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Operation Data can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	for id, depoData := range input.DepoData {

		if depoData.Asset == 0 {
			var info InfoMessage
			info.Code = 302
			info.Message = "Asset in row " + strconv.Itoa(id+1) + " can't be empty"
			json.NewEncoder(w).Encode(info)
			return
		}

		if depoData.Account == 0 {
			var info InfoMessage
			info.Code = 302
			info.Message = "Account in row " + strconv.Itoa(id+1) + " can't be empty"
			json.NewEncoder(w).Encode(info)
			return
		}

		if depoData.Count == 0 {
			var info InfoMessage
			info.Code = 302
			info.Message = "Asset Count in row " + strconv.Itoa(id+1) + " can't be empty"
			json.NewEncoder(w).Encode(info)
			return
		}

	}

	operationID, status := DepoOperationNewItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "New Depo Operation with ID = " + strconv.Itoa(operationID) + " added"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	}

}

func UpdateInOutDepo(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func GetInOutDepo(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func DeleteInOutDepo(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func NewInOutBank(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func UpdateInOutBank(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func GetInOutBank(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}

func DeleteInOutBank(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

}
