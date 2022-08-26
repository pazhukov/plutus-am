package plutus

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Currency struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type CurrencyRate struct {
	Period   string  `json:"period"`
	Currency int     `json:"currency"`
	Rate     float64 `json:"rate"`
}

func NewCurrency(w http.ResponseWriter, r *http.Request) {

	var input Currency
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Title == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Currency Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.ID == 0 {
		var info InfoMessage
		info.Code = 302
		info.Message = "Currency ID can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	// add to db
	var status = CurrencyNewItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "New Currency added"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == -2 {
		var info InfoMessage
		info.Code = 303
		info.Message = "Currency with code " + strconv.Itoa(input.ID) + " exist"
		json.NewEncoder(w).Encode(info)
	}

}

func UpdateCurrency(w http.ResponseWriter, r *http.Request) {

	var input Currency
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Title == "" {
		var info InfoMessage
		info.Code = 301
		info.Message = "Currency Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.ID == 0 {
		var info InfoMessage
		info.Code = 302
		info.Message = "Currency ID can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	// add to db
	var status = CurrencyUpdateItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "Currency updated"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == -2 {
		var info InfoMessage
		info.Code = 303
		info.Message = "Currency with code " + strconv.Itoa(input.ID) + " not found"
		json.NewEncoder(w).Encode(info)
	}

}

func GetCurrency(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	currencyID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert currency ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	currency, status := GetCurrencyByID(currencyID)
	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 300
		info.Message = "Currency with ID = " + strconv.Itoa(currencyID) + " not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		json.NewEncoder(w).Encode(currency)
	}

}

func DeleteCurrency(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	currencyID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert currency ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	status := DeleteCurrencyByID(currencyID)
	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 300
		info.Message = "Currency with ID = " + strconv.Itoa(currencyID) + " not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "Currency deleted"
		json.NewEncoder(w).Encode(info)
	}

}

func LoadCurrencyRates(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	data := GetCurrencyRatesCBR("dfsdf")

	log.Println(data)

}
