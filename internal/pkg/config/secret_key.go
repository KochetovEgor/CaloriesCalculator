package config

import (
	"fmt"
	"io"
	"os"
)

func LoadSecretKey(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", path, err)
	}
	defer file.Close()

	secretKey, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", path, err)
	}
	return secretKey, nil
}
