package plutus

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func openDBConnection() {
	db, err := sql.Open("mysql", "root:usbw@tcp(localhost:3307)/plutus-am?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	database = db

}

func closeDBConnection() {
	database.Close()

}

func AssetNewItem(item Asset) int {
	openDBConnection()

	var lastInsertId int64

	res, err := database.Exec("INSERT INTO assets (title, type_id, currency_id) VALUES (?, ?, ?)", item.Title, item.TypeID, item.CurrencyID)
	if err != nil {
		log.Println("Insert Asset")
		log.Println(err)
		closeDBConnection()
		return -1
	}

	lastInsertId, err = res.LastInsertId()
	if err != nil {
		log.Println("Get Last ID")
		log.Println(err)
		closeDBConnection()
		return -1
	}

	for _, element := range item.Codes {
		_, err := database.Exec("INSERT INTO assets_codes (asset, code) VALUES (?, ?)", lastInsertId, element.Code)
		if err != nil {
			log.Println("Insert Codes")
			log.Println(err)
			closeDBConnection()
			return -1
		}
	}

	closeDBConnection()

	return 200

}

func GetAssetByID(assetID int) (Asset, int) {
	openDBConnection()

	asset := Asset{}
	codes := []AssetCode{}

	rows, err := database.Query("SELECT id, title, type_id, currency_id FROM assets WHERE id = ?", assetID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return asset, -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&asset.ID, &asset.Title, &asset.TypeID, &asset.CurrencyID)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return asset, -1
		}
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return asset, -2
	}

	err = rows.Close()
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return asset, -1
	}

	// codes
	rows_c, err_c := database.Query("SELECT code FROM assets_codes WHERE asset = ?", assetID)
	if err_c != nil {
		log.Println(err_c)
		closeDBConnection()
		return asset, -1
	}

	defer rows_c.Close()
	for rows_c.Next() {
		code := AssetCode{}
		err := rows_c.Scan(&code.Code)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return asset, -1
		}
		codes = append(codes, code)
	}

	asset.Codes = codes

	closeDBConnection()

	return asset, 200

}

func DeleteAssetByID(assetID int) int {

	openDBConnection()

	rows, err := database.Query("SELECT id FROM assets WHERE id = ?", assetID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return 404

	}

	_, err_del1 := database.Exec("DELETE from assets WHERE id = ?", assetID)
	if err_del1 != nil {
		log.Println("Delete Asset")
		log.Println(err_del1)
		closeDBConnection()
		return -1
	}

	_, err_del2 := database.Exec("DELETE from assets_codes WHERE asset = ?", assetID)
	if err_del2 != nil {
		log.Println("Delete Code")
		log.Println(err_del2)
		closeDBConnection()
		return -1
	}

	closeDBConnection()

	return 200

}

func UpdateAssetInDB(asset Asset) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM assets WHERE id = ?", asset.ID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return 404

	}

	_, err_upd := database.Exec("UPDATE assets SET title = ?, type_id = ?, currency_id = ? WHERE id = ?", asset.Title, asset.TypeID, asset.CurrencyID, asset.ID)
	if err_upd != nil {
		log.Println("Update Asset")
		log.Println(err_upd)
		closeDBConnection()
		return -1
	}

	_, err_del1 := database.Exec("DELETE from assets_codes WHERE asset = ?", asset.ID)
	if err_del1 != nil {
		log.Println("Delete Asset Code")
		log.Println(err_del1)
		closeDBConnection()
		return -1
	}

	for _, element := range asset.Codes {
		_, err := database.Exec("INSERT INTO assets_codes (asset, code) VALUES (?, ?)", asset.ID, element.Code)
		if err != nil {
			log.Println("Insert Codes")
			log.Println(err)
			closeDBConnection()
			return -1
		}
	}

	closeDBConnection()

	return 200

}

func AddNewAssetCode(item AssetCodeHTTP) int {
	openDBConnection()

	rows, err := database.Query("SELECT id, title, type_id, currency_id FROM assets WHERE id = ?", item.ID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return 404

	}

	rows_c, err_c := database.Query("SELECT asset, code FROM assets_codes WHERE code = ?", item.Code)
	if err_c != nil {
		log.Println(err_c)
		closeDBConnection()
		return -1
	}

	count_c := 0
	defer rows_c.Close()
	for rows_c.Next() {
		count_c = count_c + 1
	}

	if count_c > 0 {
		closeDBConnection()
		return 300

	}

	_, err_i := database.Exec("INSERT INTO assets_codes (asset, code) VALUES (?, ?)", item.ID, item.Code)
	if err_i != nil {
		log.Println("Insert Codes")
		log.Println(err_i)
		closeDBConnection()
		return -1
	}

	closeDBConnection()

	return 200

}

func DeleteAssetCode(code AssetCodeHTTP) int {

	openDBConnection()

	rows, err := database.Query("SELECT asset, code  FROM assets_codes WHERE asset = ? AND code = ?", code.ID, code.Code)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return 404

	}

	_, err_i := database.Exec("DELETE  FROM assets_codes WHERE asset = ? AND code = ?", code.ID, code.Code)
	if err_i != nil {
		log.Println("Delete Code")
		log.Println(err_i)
		closeDBConnection()
		return -1
	}

	closeDBConnection()

	return 200

}

func GetAssetList(page int) ([]Asset, int) {
	openDBConnection()

	var assets []Asset

	rows, err := database.Query("SELECT id, title, type_id, currency_id FROM assets WHERE id > ? ORDER BY id LIMIT 10", page)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return assets, -1
	}

	current_page := 0
	defer rows.Close()
	for rows.Next() {
		asset := Asset{}
		err := rows.Scan(&asset.ID, &asset.Title, &asset.TypeID, &asset.CurrencyID)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return assets, -1
		}

		codes := []AssetCode{}

		rows_c, err_c := database.Query("SELECT code FROM assets_codes WHERE asset = ?", asset.ID)
		if err_c != nil {
			log.Println(err_c)
			closeDBConnection()
			return assets, -1
		}

		defer rows_c.Close()
		for rows_c.Next() {
			code := AssetCode{}
			err := rows_c.Scan(&code.Code)
			if err != nil {
				log.Println(err)
				closeDBConnection()
				return assets, -1
			}
			codes = append(codes, code)
		}

		current_page = asset.ID
		asset.Codes = codes
		assets = append(assets, asset)
	}

	closeDBConnection()

	return assets, current_page

}

func CurrencyNewItem(item Currency) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM currencies WHERE id = ?", item.ID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		count = count + 1
	}

	if count > 0 {
		closeDBConnection()
		return -2
	}
	_, err2 := database.Exec("INSERT INTO currencies (id, title) VALUES (?, ?)", item.ID, item.Title)
	if err2 != nil {
		log.Println(err2)
		closeDBConnection()
		return -1
	}

	closeDBConnection()

	return 200

}

func CurrencyUpdateItem(item Currency) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM currencies WHERE id = ?", item.ID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return -2
	}

	_, err2 := database.Exec("UPDATE currencies SET id = ?, title = ? WHERE id = ? ", item.ID, item.Title, item.ID)
	if err2 != nil {
		log.Println(err2)
		closeDBConnection()
		return -1
	}

	closeDBConnection()

	return 200

}

func GetCurrencyByID(currencyID int) (Currency, int) {
	openDBConnection()

	currency := Currency{}

	rows, err := database.Query("SELECT id, title FROM currencies WHERE id = ?", currencyID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return currency, -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&currency.ID, &currency.Title)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return currency, -1
		}
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return currency, 404
	}

	err = rows.Close()
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return currency, -1
	}

	closeDBConnection()

	return currency, 200

}

func DeleteCurrencyByID(currencyID int) int {
	openDBConnection()

	rows, err := database.Query("SELECT id, title FROM currencies WHERE id = ?", currencyID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return 404
	}

	err = rows.Close()
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return -1
	}

	_, err_del1 := database.Exec("DELETE from currencies WHERE id = ?", currencyID)
	if err_del1 != nil {
		log.Println(err_del1)
		closeDBConnection()
		return -1
	}

	closeDBConnection()

	return 200

}

func LoadCurrencyRateInDB(rates []CurrencyRate) (int, int) {
	openDBConnection()

	var inserted int
	for _, element := range rates {

		rows, err := database.Query("SELECT period FROM currency_rates WHERE period = ? and currency_id = ?", element.Period, element.Currency)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return 0, -1
		}

		count := 0
		defer rows.Close()
		for rows.Next() {
			count = count + 1
		}

		if count > 0 {

			_, err_del := database.Exec("DELETE from currency_rates WHERE period = ? and currency_id = ?", element.Period, element.Currency)
			if err_del != nil {
				log.Println(err_del)
				closeDBConnection()
				return 0, -1
			}

		}

		_, err1 := database.Exec("INSERT INTO currency_rates (period, currency_id, rate) VALUES (?, ?, ?)", element.Period, element.Currency, element.Rate)
		if err1 != nil {
			log.Println(err1)
			closeDBConnection()
			return 0, -1
		}

		inserted = inserted + 1

	}

	closeDBConnection()

	return inserted, 200

}

func GetCurrencyRatesCBR(onDate string) ([]CurrencyRate, int) {
	openDBConnection()

	rates := []CurrencyRate{}

	var sql string
	if onDate == "1900-01-01" {
		sql = "SELECT max_rates.period as period, max_rates.currency_id as currency_id, IFNULL(rates.rate, 0) as rate FROM (SELECT MAX(period) as period, currency_id FROM currency_rates GROUP BY currency_id) as max_rates LEFT JOIN currency_rates as rates ON max_rates.period = rates.period AND max_rates.currency_id = rates.currency_id"
	} else {
		sql = "SELECT max_rates.period as period, max_rates.currency_id as currency_id, IFNULL(rates.rate, 0) as rate FROM (SELECT MAX(period) as period, currency_id FROM currency_rates WHERE period <= '" + onDate + "' GROUP BY currency_id) as max_rates LEFT JOIN currency_rates as rates ON max_rates.period = rates.period AND max_rates.currency_id = rates.currency_id"
	}

	rows, err := database.Query(sql)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return rates, -1
	}

	defer rows.Close()
	for rows.Next() {
		rate := CurrencyRate{}
		err := rows.Scan(&rate.Period, &rate.Currency, &rate.Rate)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return rates, -1
		}
		rates = append(rates, rate)
	}

	closeDBConnection()

	return rates, 200

}
