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

func Write(collectionName string, data interface{}) error {
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

func Update(collectionName string, data interface{}) error {
	_path := collectionName
	_pathSwap := fmt.Sprintf("%s~swap", _path)
	_pathOld := fmt.Sprintf("%s~old", _path)
	var err error

	defer func() {
		_pathSwap = path(_pathSwap)
		_path = path(_path)
		_pathOld = path(_pathOld)

		if err != nil {
			os.Remove(_pathSwap)
		} else {
			//all done
			//back up
			os.Rename(_path, _pathOld)

			//rename ~swap to path
			err := os.Rename(_pathSwap, _path)
			if err == nil {
				//remove the old data
				os.Remove(_pathOld)
				os.Remove(_pathSwap)
			}

		}

	}()

	err = Write(_pathSwap, data)
	if err != nil {
		return err
	}

	return nil
}
