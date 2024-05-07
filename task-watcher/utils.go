package watcher

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func createPath(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func readResultFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func writeResultFile(path string, data []byte) error {
	err := createPath(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
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

type GPUInfo struct {
	Name   string
	Memory uint64
}

func getGPUInfo() (*GPUInfo, error) {
	// nvidia-smi --query-gpu=gpu_name,memory.total --format=csv,noheader

	output, err := exec.Command("nvidia-smi", "--query-gpu=gpu_name,memory.total", "--format=csv,noheader").Output()
	if err != nil {
		return nil, err
	}

	// output example:
	// GeForce RTX 2080 Ti, 11019 MiB

	// split by comma
	parts := strings.Split(string(output), ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid output")
	}

	// parse memory
	var memory uint64
	_, err = fmt.Sscanf(parts[1], "%d MiB", &memory)
	if err != nil {
		return nil, err
	}

	return &GPUInfo{
		Name:   strings.TrimSpace(parts[0]),
		Memory: memory,
	}, nil

}
