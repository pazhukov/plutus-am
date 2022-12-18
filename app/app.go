package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"plutus"
)

var pathToConfig string

func main() {

	flag.StringVar(&pathToConfig, "dbcfg", "db.json", "path to db config")
	flag.Parse()

	log.Println("cfg:", pathToConfig)

	file, err_file := os.Open(pathToConfig)
	if err_file != nil {
		log.Fatal(err_file)
	}
	decoder := json.NewDecoder(file)
	dbcfg := plutus.ConnectionSetting{}
	err := decoder.Decode(&dbcfg)
	if err != nil {
		log.Println("error:", err)
	}

	plutus.ConnectionString = dbcfg.User + ":" + dbcfg.Password + "@tcp"
	plutus.ConnectionString = plutus.ConnectionString + "(" + dbcfg.Server + ":" + dbcfg.Port + ")/" + dbcfg.DBname
	plutus.ConnectionString = plutus.ConnectionString + "?charset=" + dbcfg.Charset

	plutus.StartServer()

}
