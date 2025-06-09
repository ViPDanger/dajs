package db

import (
	"encoding/json"
	"os"
	"regexp"
)

func Read(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath) // Замените на ваш путь
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`(?s)/\*.*?\*/`)
	return re.ReplaceAll(data, []byte("")), nil
}

func Compile[T any](filepath string) (object T, err error) {
	data, err := Read(filepath)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &object)
	return
}
