package db

import (
	"log"
	"github.com/dgraph-io/badger/v4"
)

var DB *badger.DB

func InitBadger() {
	opts := badger.DefaultOptions("./data")
	opts.Logger = nil // disable verbose logs
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Failed to open BadgerDB: %v", err)
	}
	DB = db
}

func CloseBadger() {
	DB.Close()
}
