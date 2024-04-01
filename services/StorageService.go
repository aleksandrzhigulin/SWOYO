package services

import (
	"SWOYO/db"
	"SWOYO/storage"
)

var databaseConnected bool
var dbInstance db.Database

func StorageServiceInit(database db.Database, databaseMode bool) {
	databaseConnected = databaseMode
	if databaseMode {
		dbInstance = database
	}
}

func SetToStorage(key string, value string) {
	if databaseConnected {
		err := dbInstance.SaveShortUrl(key, value)
		if err != nil {
			return
		}
	} else {
		storage.Storage[key] = value
	}
}

func ItemExists(key string) bool {
	if databaseConnected {
		return dbInstance.ShortUrlExists(key)
	}
	return len(storage.Storage[key]) > 0
}

func GetFromStorage(key string) string {
	if databaseConnected {
		_, result := dbInstance.GetFullUrl(key)
		return result
	}
	return storage.Storage[key]
}
