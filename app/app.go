package main

import (
	"log"
	"os"
	"path/filepath"
	"plutus"
)

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)

	plutus.StartServer()

}
