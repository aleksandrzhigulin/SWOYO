package storage

var Storage map[string]string

func StorageInit() {
	Storage = make(map[string]string)
}
