package manager

import (
	"math/rand"
	"os"
	"path/filepath"
)

func createPath(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func getCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func randPort() int {
	rnd := rand.Intn(1000) + 10000
	return rnd
}
