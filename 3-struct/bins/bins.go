package bins

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

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id,
		private,
		createdAt,
		name,
	}
}

func NewBinList() *BinList {
	return &BinList{}
}
