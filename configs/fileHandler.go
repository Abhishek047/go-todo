package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	apptype "github.com/Abhishek047/todo-app/app-type"
)

type FileData interface{}

// @open a file with path
func OpenFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fmt.Println("open file...")
	return io.ReadAll(file)
}

// @create default config file
func CreateDefaultConfig() (*apptype.AppConfig, error) {
	fmt.Println("creating directory...")
	err := os.MkdirAll(apptype.ConfigDirectory, 0700)
	if err != nil {
		return nil, err
	}
	fmt.Println("creating file...")
	newFile, err := os.Create(apptype.ConfigPath)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()
	fmt.Println("Encoding...")
	encoder := json.NewEncoder(newFile)
	err = encoder.Encode(DefaultConfig)
	if err != nil {
		return nil, err
	}
	return &DefaultConfig, nil
}

// open config file
func GetConfigData() (*apptype.AppConfig, error) {
	fmt.Println("fetching config")
	data, err := OpenFile(apptype.ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			config, err := CreateDefaultConfig()
			if err != nil {
				return nil, errors.New(fmt.Sprint("Error in CreateDefaultConfig error: ", err))
			}
			return config, nil
		}
		return nil, err
	}
	// can make this un-marshalling as a strategy
	var typedData apptype.AppConfig
	err = json.Unmarshal(data, &typedData)
	if err != nil {
		return nil, err
	}
	return &typedData, nil
}
