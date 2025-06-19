package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id,
		private,
		createdAt,
		name,
	}
}

func main() {
	_ = newBin("", false, time.Now(), "")
}
