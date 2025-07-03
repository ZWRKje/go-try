package main

import (
	"fmt"
	"struct/api"
	"struct/bins"
	"struct/config"
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
	cfg := config.NewConfig()
	api := api.NewApi(*cfg)
	fmt.Print(api)
}
