package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/files"
)

func ReadConfigFromFile[T any](filePath string) (*T, error) {
	if !files.IsFileExist(filePath) {
		return nil, errors.New(fmt.Sprintf("file '%s' not found", filePath))
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Join(err, errors.New(fmt.Sprintf("cannot read file '%v'", filePath)))
	}

	cfg := new(T)
	if err := json.Unmarshal(fileBytes, cfg); err != nil {
		return nil, errors.Join(err, errors.New(fmt.Sprintf("cannot unmarshal config file '%v'", filePath)))
	}

	return cfg, nil
}
