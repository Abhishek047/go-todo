package main

import (
	"fmt"

	"github.com/Abhishek047/todo-app/configs"
)

func main() {
	fmt.Println("main")
	var config configs.AppConfig
	err := configs.OpenFile("config12.json", &config)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config)
	}
}
