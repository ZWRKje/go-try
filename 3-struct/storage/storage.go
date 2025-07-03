package storage

import (
	"encoding/json"
	"struct/bins"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type IBinList interface {
	AddBin(bins.Bin)
}

type Storage struct {
	Bins      IBinList  `json:"bins"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type StorageWithDb struct {
	Storage
	db Db
}

func NewStorage(db Db, bins IBinList) *StorageWithDb {
	file, err := db.Read()
	if err != nil {
		return &StorageWithDb{
			Storage: Storage{
				Bins:      bins,
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var storage Storage

	err = json.Unmarshal(file, &storage)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &StorageWithDb{
			Storage: Storage{
				Bins:      bins,
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	return &StorageWithDb{
		Storage: storage,
		db:      db,
	}
}

func (st *StorageWithDb) ToBytes() ([]byte, error) {
	file, err := json.Marshal(st.Storage)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (st *StorageWithDb) SaveInfo() {
	data, err := st.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}

	st.db.Write(data)
}
