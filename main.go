package main

import (
	"fmt"

	"github.com/Abhishek047/todo-app/configs"
)

func main() {
	fmt.Println("main")
	data, err := configs.GetConfigData()
	if err != nil {
		fmt.Println("we have an error")
	}
	fmt.Println(data)
}
