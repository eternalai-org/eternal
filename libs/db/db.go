package db

import (
	"fmt"
	"github.com/goccy/go-json"
	"os"
)

const DB_PATH = "./database"

func path(collectionName string) string {
	return fmt.Sprintf("%s/%s.json", DB_PATH, collectionName)
}

func read(collectionName string) ([]byte, error) {
	_path := path(collectionName)

	_byte, err := os.ReadFile(_path)
	if err != nil {
		return nil, err
	}

	return _byte, nil
}

func Read(collectionName string, out interface{}) error {

	_b, err := read(collectionName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(_b, &out)
	if err != nil {
		return err
	}

	return nil
}

func Write(data interface{}, collectionName string) error {
	_path := path(collectionName)
	// Check if the folder exists
	if _, err := os.Stat(DB_PATH); os.IsNotExist(err) {
		err = os.Mkdir(DB_PATH, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(_path, _b, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
