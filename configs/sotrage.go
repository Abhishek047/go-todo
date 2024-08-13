package configs

import (
	apptype "github.com/Abhishek047/todo-app/app-type"
)

var DefaultConfig apptype.AppConfig = apptype.AppConfig{
	Db:             apptype.Device,
	DefaultStorage: apptype.Device,
	Path:           apptype.StoragePath,
}
