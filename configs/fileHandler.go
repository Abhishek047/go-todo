package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileData interface{}

func OpenFile(path string, typedData interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("no-file create one")
		} else {
			return err
		}
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(typedData)
	if err != nil {
		return err
	}
	return nil
}
