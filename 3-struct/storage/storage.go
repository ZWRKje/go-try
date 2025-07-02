package storage

import (
	"encoding/json"
	"fmt"
	"struct/bins"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Storage struct {
	Bins      *bins.BinList `json:"bins"`
	UpdatedAt time.Time     `json:"updatedAt"`
	db        Db
}

func NewStorage(db Db) *Storage {
	file, err := db.Read()
	if err != nil {
		return &Storage{
			Bins:      bins.NewBinList(),
			UpdatedAt: time.Now(),
			db:        db,
		}
	}

	var storage Storage

	err = json.Unmarshal(file, &storage)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
	}

	return &storage
}

func (st *Storage) ToBytes() ([]byte, error) {
	file, err := json.Marshal(st)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (st *Storage) SaveInfo() {
	data, err := st.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	fmt.Println(data)
	st.db.Write(data)
}
