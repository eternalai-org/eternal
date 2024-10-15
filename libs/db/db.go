package db

import (
	"eternal-infer-worker/model_structures"
	"fmt"
	"github.com/goccy/go-json"
	"os"
)

const DB_PATH = "./database"

func path(collectionName string) string {
	return fmt.Sprintf("%s/%s.json", DB_PATH, collectionName)
}

func Read(collectionName string) ([]byte, error) {
	_path := path(collectionName)

	_byte, err := os.ReadFile(_path)
	if err != nil {
		return nil, err
	}

	return _byte, nil
}

func Query() (*model_structures.IModel, error) {

	return nil, nil
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
