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

	res, err := database.Exec("insert into assets (title, type_id, currency_id) values (?, ?, ?)", item.Title, item.TypeID, item.CurrencyID)
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
		_, err := database.Exec("insert into assets_codes (asset, code) values (?, ?)", lastInsertId, element.Code)
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

	if count > 0 {
		closeDBConnection()
		return 300

	}

	_, err_i := database.Exec("insert into assets_codes (asset, code) values (?, ?)", item.ID, item.Code)
	if err_i != nil {
		log.Println("Insert Codes")
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

	rows, err := database.Query("SELECT id, title, type_id, currency_id FROM assets WHERE id > ? order by id limit 10", page)
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
