package main

import (
	"struct/bins"
	"struct/file"
	"struct/storage"
	"time"
)

func main() {
	bin := bins.NewBin("", false, time.Now(), "")
	fileDb := file.NewFileDb("data.json")
	storage := storage.NewStorage(fileDb)
	storage.Bins.AddBin(*bin)
	storage.SaveInfo()
}
