package main

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	bins []Bin
}

func newBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id,
		private,
		createdAt,
		name,
	}
}

func newBinList() *BinList {
	return &BinList{}
}

func main() {
	_ = newBin("", false, time.Now(), "")
}
