package database

import (
	"encoding/json"
	"fmt"
	"os"

	apptype "github.com/Abhishek047/todo-app/app-type"
	"github.com/Abhishek047/todo-app/configs"
)

type deviceStart struct{}

// My object pool
var DeviceStrategy = &deviceStart{}
var list = []apptype.Todo{}
var settings = configs.GetConfig()

// My object pool

func getList() (*[]apptype.Todo, error) {
	data, err := configs.OpenFile(settings.Path)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		list = []apptype.Todo{}
		return nil, nil
	}
	fmt.Println("Decoding todo...")
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (s *deviceStart) Save(item apptype.Todo) error {
	_, err := getList()
	if err != nil {
		return err
	}
	list = append(list, item)
	result, err := json.Marshal(list)
	if err != nil {
		return err
	}
	os.WriteFile(settings.Path, result, os.ModePerm)
	return nil
}
func (s *deviceStart) Fetch() ([]apptype.Todo, error) {
	// fetch the list from file
	_, err := getList()
	if err != nil {
		// if not found create todo
		if os.IsNotExist(err) {
			fmt.Println("creating todo file...")
			newFile, err := os.Create(settings.Path)
			if err != nil {
				return nil, err
			}
			defer newFile.Close()
		} else {
			panic(err)
		}
	}
	// return the global list from here
	return list, nil
}
