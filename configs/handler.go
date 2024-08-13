package configs

import (
	"fmt"
	"sync"

	apptype "github.com/Abhishek047/todo-app/app-type"
)

var lock = &sync.Mutex{}

var config *apptype.AppConfig

func GetConfig() *apptype.AppConfig {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
		if config == nil {
			data, err := GetConfigData()
			if err != nil {
				fmt.Println("err in fetching config....")
				config = &apptype.AppConfig{
					Db:             apptype.Device,
					DefaultStorage: apptype.Device,
					Path:           apptype.StoragePath,
				}
			} else {
				config = data
			}
		}
	} else {
		fmt.Println("Already fetched")
	}
	return config
}
