package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	apptype "github.com/Abhishek047/todo-app/app-type"
	"github.com/Abhishek047/todo-app/configs"
)

type deviceStart struct {
	name apptype.DBType
}

// My object pool
var DeviceStrategy = &deviceStart{
	name: apptype.Device,
}
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
	// fmt.Println("Decoding todo...")
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (s *deviceStart) AddItem(item apptype.Todo) error {
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
			// fmt.Println("creating todo file...")
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
func (s *deviceStart) GetName() apptype.DBType {
	return s.name
}

func (s *deviceStart) UpdateTodoDone(id string) *apptype.Todo {
	_, err := s.Fetch()
	if err != nil {
		log.Fatal(err)
	}
	idForSearch, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	var found int = -1
	for index, item := range list {
		if item.Id == int64(idForSearch) {
			found = index
			break
		}
	}
	if found == -1 {
		fmt.Println("Not found")
		return nil
	}
	list[found].Status = true

	foundItem := list[found]
	result, err := json.Marshal(list)
	if err != nil {
		log.Fatal("error in marshaling")
		return nil
	}
	os.WriteFile(settings.Path, result, os.ModePerm)
	return &foundItem
}
