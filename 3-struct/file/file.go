package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	ext := filepath.Ext(name)

	if strings.ToLower(ext) != ".json" {
		return nil, errors.New("File extension must be json")
	}

	file, err := os.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return file, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
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
