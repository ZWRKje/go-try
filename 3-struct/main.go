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
	binList := bins.NewBinList()
	storage := storage.NewStorage(fileDb, binList)
	storage.Bins.AddBin(*bin)
	storage.SaveInfo()
}
