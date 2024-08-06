package configs

import (
	"fmt"
	"sync"
)

type AppConfig struct {
	Db             DBType `json:"db"`
	DefaultStorage DBType `json:"defaultStorage"`
	Path           string `json:"path"`
}

var lock = &sync.Mutex{}

var config *AppConfig

func GetConfig() *AppConfig {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
		if config == nil {
			// fetch configs
			config = &AppConfig{
				Db:             Device,
				DefaultStorage: Device,
				Path:           StoragePath,
			}
		}
	} else {
		fmt.Println("Already fetched")
	}
	return config
}
