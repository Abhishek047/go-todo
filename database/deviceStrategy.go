package database

import (
	"fmt"

	"github.com/Abhishek047/todo-app/todo"
)

type deviceStart struct{}

var DeviceStrategy = &deviceStart{}

func (s *deviceStart) Save(item todo.Todo) {
	fmt.Println(item)
}
func (s *deviceStart) Fetch() (*todo.TodoList, error) {
	// data, err := configs.OpenFile(configs.StoragePath)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// fmt.Println("creating todo file...")
	// newFile, err := os.Create(ConfigPath)
	// if err != nil {
	// 	return nil, err
	// }
	// defer newFile.Close()
	// fmt.Println("Encoding...")
	// encoder := json.NewEncoder(newFile)
	// err = encoder.Encode(DefaultConfig)
	// if err != nil {
	// 	return nil, err
	// }
	// } else {
	// 	panic(err)
	// }
	// }
	return &todo.TodoList{
		List: []todo.Todo{},
	}, nil
}
