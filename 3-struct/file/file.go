package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileDb struct {
	name string
}

func NewFileDb(fileName string) *FileDb {
	return &FileDb{name: fileName}
}

func (db *FileDb) Read() ([]byte, error) {
	ext := filepath.Ext(db.name)

	if strings.ToLower(ext) != ".json" {
		return nil, errors.New("File extension must be json")
	}

	file, err := os.ReadFile(db.name)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return file, nil
}

func (db *FileDb) Write(content []byte) {
	file, err := os.Create(db.name)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(content)
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")
	file.Close()
}
