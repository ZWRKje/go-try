package storage

import (
	"encoding/json"
	"struct/bins"
	"struct/file"
	"time"

	"github.com/fatih/color"
)

type Storage struct {
	Bins      bins.BinList `json:"bins"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

func NewStorage() *Storage {
	file, err := file.ReadFile("data.json")
	if err != nil {
		return &Storage{
			Bins:      *bins.NewBinList(),
			UpdatedAt: time.Now(),
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

	file.WriteFile(data, "data.json")
}
