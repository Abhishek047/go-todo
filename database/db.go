package database

import (
	"fmt"
	"sync"

	"github.com/Abhishek047/todo-app/configs"
	"github.com/Abhishek047/todo-app/todo"
)

type DbI interface {
	Save(todo.Todo)
	Fetch() todo.TodoList
}
type DbClass struct {
	DB DbI
}

var single = &DbClass{}
var lock = &sync.Mutex{}

func setDb(DbStrategy DbI) {
	single.DB = DbStrategy
}

func GetDb() DbI {
	if single.DB == nil {
		lock.Lock()
		defer lock.Unlock()
		if single.DB == nil {
			data := configs.GetConfig()
			fmt.Println(data, "current")
			switch data.Db {
			case configs.Mongo:
				{
					fmt.Println("set To mongo")
					break
				}
			case configs.Device:
				{
					setDb(DeviceStrategy)
					break
				}
			case configs.Psql:
				{
					fmt.Println("set To psql")
					break
				}
			}
		}
	} else {
		fmt.Println("Already fetched")
	}
	return single.DB
}
