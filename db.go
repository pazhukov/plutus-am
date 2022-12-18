package plutus

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB
var ConnectionString string

type ConnectionSetting struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
	DBname   string `json:"dbname"`
	Charset  string `json:"charset"`
}

func openDBConnection() {
	db, err := sql.Open("mysql", ConnectionString)
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

func OwnerNewItem(item Owner) (int, int) {
	openDBConnection()

	var ownerID int
	var lastInsertId int64

	res, err := database.Exec("INSERT INTO owners (title) VALUES (?)", item.Title)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return ownerID, -1
	}

	lastInsertId, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return ownerID, -1
	}

	ownerID = int(lastInsertId)

	closeDBConnection()
	return ownerID, 200

}

func OnwerUpdateItem(item Owner) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM owners WHERE id = ?", item.ID)
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

	_, err_upd := database.Exec("UPDATE owners SET title = ? WHERE id = ?", item.Title, item.ID)
	if err_upd != nil {
		log.Println(err_upd)
		closeDBConnection()
		return -1
	}

	closeDBConnection()
	return 200

}

func OnwerGetItem(ownerID int) (Owner, int) {
	openDBConnection()

	item := Owner{}

	rows, err := database.Query("SELECT id, title FROM owners WHERE id = ?", ownerID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return item, -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Title)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return item, -1
		}
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return item, 404
	}

	closeDBConnection()
	return item, 200

}

func OnwerDeleteItem(ownerID int) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM owners WHERE id = ?", ownerID)
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

	_, err_upd := database.Exec("DELETE FROM owners WHERE id = ?", ownerID)
	if err_upd != nil {
		log.Println(err_upd)
		closeDBConnection()
		return -1
	}

	closeDBConnection()
	return 200

}

func OwnerListItem() ([]Owner, int) {
	openDBConnection()

	list := []Owner{}

	rows, err := database.Query("SELECT id, title FROM owners")
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return list, -1
	}

	defer rows.Close()
	for rows.Next() {
		item := Owner{}
		err := rows.Scan(&item.ID, &item.Title)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return list, -1
		}
		list = append(list, item)
	}

	closeDBConnection()
	return list, 200

}

func DepoAccountNewItem(item DepoAccount) (int, int) {

	openDBConnection()

	var depoAccountID int
	var lastInsertId int64

	res, err := database.Exec("INSERT INTO depo_accounts (owner_id, title, code, broker) VALUES (?, ?, ?, ?)", item.Owner, item.Title, item.Code, item.Broker)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return depoAccountID, -1
	}

	lastInsertId, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return depoAccountID, -1
	}

	depoAccountID = int(lastInsertId)

	closeDBConnection()
	return depoAccountID, 200

}

func DepoAccountUpdateItem(item DepoAccount) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM depo_accounts WHERE id = ?", item.ID)
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

	_, err_upd := database.Exec("UPDATE depo_accounts SET owner_id = ?, title = ?, code = ?, broker = ? WHERE id = ?", item.Owner, item.Title, item.Code, item.Broker, item.ID)
	if err_upd != nil {
		log.Println(err_upd)
		closeDBConnection()
		return -1
	}

	closeDBConnection()
	return 200

}

func DepoAccountGetItem(depoAccountID int) (DepoAccount, int) {

	openDBConnection()

	item := DepoAccount{}

	rows, err := database.Query("SELECT id, owner_id, title, code, broker FROM depo_accounts WHERE id = ?", depoAccountID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return item, -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Owner, &item.Title, &item.Code, &item.Broker)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return item, -1
		}
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return item, 404
	}

	closeDBConnection()
	return item, 200

}

func DepoAccountDeleteItem(DepoAccount int) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM depo_accounts WHERE id = ?", DepoAccount)
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

	_, err_upd := database.Exec("DELETE FROM depo_accounts WHERE id = ?", DepoAccount)
	if err_upd != nil {
		log.Println(err_upd)
		closeDBConnection()
		return -1
	}

	closeDBConnection()
	return 200

}

func BankAccountNewItem(item BankAccount) (int, int) {

	openDBConnection()

	var depoAccountID int
	var lastInsertId int64

	res, err := database.Exec("INSERT INTO bank_accounts (owner_id, title, currency_id, code, broker) VALUES (?, ?, ?, ?, ?)", item.Owner, item.Title, item.Currency, item.Code, item.Broker)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return depoAccountID, -1
	}

	lastInsertId, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return depoAccountID, -1
	}

	depoAccountID = int(lastInsertId)

	closeDBConnection()
	return depoAccountID, 200

}

func BankAccountUpdateItem(item BankAccount) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM bank_accounts WHERE id = ?", item.ID)
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

	_, err_upd := database.Exec("UPDATE bank_accounts SET owner_id = ?, title = ?, currency_id = ?, code = ?, broker = ? WHERE id = ?", item.Owner, item.Title, item.Currency, item.Code, item.Broker, item.ID)
	if err_upd != nil {
		log.Println(err_upd)
		closeDBConnection()
		return -1
	}

	closeDBConnection()
	return 200

}

func BankAccountGetItem(bankAccountID int) (BankAccount, int) {

	openDBConnection()

	item := BankAccount{}

	rows, err := database.Query("SELECT id, owner_id, title, currency_id, code, broker FROM bank_accounts WHERE id = ?", bankAccountID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return item, -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Owner, &item.Title, &item.Currency, &item.Code, &item.Broker)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return item, -1
		}
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return item, 404
	}

	closeDBConnection()
	return item, 200

}

func BankAccountDeleteItem(bankAccountID int) int {
	openDBConnection()

	rows, err := database.Query("SELECT id FROM bank_accounts WHERE id = ?", bankAccountID)
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

	_, err_upd := database.Exec("DELETE FROM bank_accounts WHERE id = ?", bankAccountID)
	if err_upd != nil {
		log.Println(err_upd)
		closeDBConnection()
		return -1
	}

	closeDBConnection()
	return 200

}

func OwnerGetAccounts(ownerID int) (OwnerAccounts, int) {
	openDBConnection()

	item := OwnerAccounts{}

	rows, err := database.Query("SELECT id FROM owners WHERE id = ?", ownerID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return item, -1
	}

	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&item.Owner)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return item, -1
		}
		count = count + 1
	}

	if count == 0 {
		closeDBConnection()
		return item, 404
	}

	// Depo Account List
	list_da := []DepoAccount{}

	rows_da, err := database.Query("SELECT id, owner_id, title, code, broker FROM depo_accounts WHERE owner_id = ?", ownerID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return item, -1
	}

	defer rows_da.Close()
	for rows_da.Next() {
		item_da := DepoAccount{}
		err := rows_da.Scan(&item_da.ID, &item_da.Owner, &item_da.Title, &item_da.Code, &item_da.Broker)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return item, -1
		}
		list_da = append(list_da, item_da)
	}

	item.Depo = list_da

	// Bank Account List
	list_ba := []BankAccount{}

	rows_ba, err := database.Query("SELECT id, owner_id, title, currency_id	, code, broker FROM bank_accounts WHERE owner_id = ?", ownerID)
	if err != nil {
		log.Println(err)
		closeDBConnection()
		return item, -1
	}

	defer rows_ba.Close()
	for rows_ba.Next() {
		item_ba := BankAccount{}
		err := rows_ba.Scan(&item_ba.ID, &item_ba.Owner, &item_ba.Title, &item_ba.Currency, &item_ba.Code, &item_ba.Broker)
		if err != nil {
			log.Println(err)
			closeDBConnection()
			return item, -1
		}
		list_ba = append(list_ba, item_ba)
	}

	item.Bank = list_ba

	closeDBConnection()
	return item, 200

}

func DepoOperationNewItem(item OperationDepoInOut) (int, int) {
	openDBConnection()

	var depoOperationID int
	//var lastInsertId int64

	closeDBConnection()
	return depoOperationID, 200
}
