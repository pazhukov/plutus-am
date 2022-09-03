package plutus

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Owner struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type OwnerList struct {
	Owners []Owner `json:"owners"`
}

func NewOwner(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	var input Owner
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
		info.Message = "Owner Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	ownerID, status := OwnerNewItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "New Owner with ID = " + strconv.Itoa(ownerID) + " added"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	}

}

func UpdateOwner(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	var input Owner
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
		info.Message = "Owner Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.ID == 0 {
		var info InfoMessage
		info.Code = 302
		info.Message = "Owner ID can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	status := OnwerUpdateItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = " Owner with ID = " + strconv.Itoa(input.ID) + " updated"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 200
		info.Message = " Owner with ID = " + strconv.Itoa(input.ID) + " not found"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	}

}

func GetOwner(w http.ResponseWriter, r *http.Request) {
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

	item, status := OnwerGetItem(ownerID)

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

func DeleteOwner(w http.ResponseWriter, r *http.Request) {
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

	status := OnwerDeleteItem(ownerID)

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
		var info InfoMessage
		info.Code = 200
		info.Message = "Owner deleted"
		json.NewEncoder(w).Encode(info)
	}

}

func GetOwnerList(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)

	list, status := OwnerListItem()

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		var oList OwnerList
		oList.Owners = list
		json.NewEncoder(w).Encode(oList)
	}

}
