package database

import (
	"fmt"
	"sync"

	apptype "github.com/Abhishek047/todo-app/app-type"
	"github.com/Abhishek047/todo-app/configs"
)

type DbClass struct {
	DB apptype.DbI
}

var single = &DbClass{}
var lock = &sync.Mutex{}

func setDb(DbStrategy apptype.DbI) {
	single.DB = DbStrategy
}

func GetDb() apptype.DbI {
	if single.DB == nil {
		lock.Lock()
		defer lock.Unlock()
		if single.DB == nil {
			data := configs.GetConfig()
			switch data.Db {
			case apptype.Mongo:
				{
					fmt.Println("set To mongo")
					break
				}
			case apptype.Device:
				{
					setDb(DeviceStrategy)
					break
				}
			case apptype.Psql:
				{
					fmt.Println("set To psql")
					break
				}
			}
		}
	}
	return single.DB
}
