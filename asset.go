package plutus

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	Share  int = 1
	Bond   int = 2
	Unit   int = 3
	DR     int = 4
	Future int = 5
	Option int = 6
)

type Asset struct {
	ID         int         `json:"id"`
	Title      string      `json:"title"`
	TypeID     int         `json:"type"`
	CurrencyID int         `json:"currency"`
	Codes      []AssetCode `json:"codes"`
}

type AssetCode struct {
	Code string `json:"code"`
}

type AssetCodeHTTP struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

type AssetList struct {
	Next_Page int     `json:"next_page"`
	Asset     []Asset `json:"assets"`
}

func NewAsset(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	var input Asset
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
		info.Message = "Asset Title can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.TypeID == 0 {
		var info InfoMessage
		info.Code = 302
		info.Message = "Asset Type can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.CurrencyID == 0 {
		var info InfoMessage
		info.Code = 303
		info.Message = "Asset Currency can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	// add to db
	var status = AssetNewItem(input)

	if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "New Asset added"
		json.NewEncoder(w).Encode(info)
	} else if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	}

}

func GetAsset(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	assetID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert asset ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	asset, status := GetAssetByID(assetID)
	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == -2 {
		var info InfoMessage
		info.Code = 300
		info.Message = "Asset not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		json.NewEncoder(w).Encode(asset)
	}

}

func DeleteAsset(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["id"]
	assetID, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert asset ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}

	status := DeleteAssetByID(assetID)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Asset with ID = " + strconv.Itoa(assetID) + " not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "Asset deleted"
		json.NewEncoder(w).Encode(info)
	}

}

func UpdateAsset(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	var input Asset
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
		info.Message = "Add Asset ID for Update"
		json.NewEncoder(w).Encode(info)
		return
	}

	var status = UpdateAssetInDB(input)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Asset with ID = " + strconv.Itoa(input.ID) + " not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "Asset updated"
		json.NewEncoder(w).Encode(info)
	}

}

func GetAssets(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	inputVar := mux.Vars(r)["page"]
	page, err := strconv.Atoi(inputVar)
	if err != nil {
		var info InfoMessage
		info.Code = 300
		info.Message = "Can't convert page ID to int"
		json.NewEncoder(w).Encode(info)
		return
	}
	assets, next_page := GetAssetList(page)

	if next_page == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else {
		var list AssetList
		list.Next_Page = next_page
		list.Asset = assets
		json.NewEncoder(w).Encode(list)
	}

}

func GetAssetsStartPage(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	assets, next_page := GetAssetList(0)

	if next_page == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else {
		var list AssetList
		list.Next_Page = next_page
		list.Asset = assets
		json.NewEncoder(w).Encode(list)
	}

}

func AddNewCode(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	var input AssetCodeHTTP
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Fatal(err.Error())
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.ID == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Asset ID can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Code == "" {
		var info InfoMessage
		info.Code = 302
		info.Message = "Asset Code can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	var status = AddNewAssetCode(input)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Asset with ID = " + strconv.Itoa(input.ID) + " not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 300 {
		var info InfoMessage
		info.Code = 300
		info.Message = "Asset code exist"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "New code added"
		json.NewEncoder(w).Encode(info)
	}

}

func DeleteCode(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	var input AssetCodeHTTP
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Fatal(err.Error())
		var info InfoMessage
		info.Code = 300
		info.Message = err.Error()
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.ID == 0 {
		var info InfoMessage
		info.Code = 301
		info.Message = "Asset ID can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	if input.Code == "" {
		var info InfoMessage
		info.Code = 302
		info.Message = "Asset Code can't be empty"
		json.NewEncoder(w).Encode(info)
		return
	}

	var status = DeleteAssetCode(input)

	if status == -1 {
		var info InfoMessage
		info.Code = 500
		info.Message = "Service error"
		json.NewEncoder(w).Encode(info)
	} else if status == 404 {
		var info InfoMessage
		info.Code = 404
		info.Message = "Asset Code not found"
		json.NewEncoder(w).Encode(info)
	} else if status == 200 {
		var info InfoMessage
		info.Code = 200
		info.Message = "Asset Code deleted"
		json.NewEncoder(w).Encode(info)
	}

}
