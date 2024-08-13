package apptype

import (
	"errors"
	"time"
)

type Todo struct {
	Id        int64      `json:"id"`
	Text      string     `json:"text"`
	Status    bool       `json:"status"`
	EndDate   *time.Time `json:"endData"`
	CreatedAt *time.Time `json:"createdAt"`
}

type TodoList struct {
	List []Todo
}

// Data base-types
type DbI interface {
	Save(Todo) error
	Fetch() ([]Todo, error)
}

// configs

type DBType string

const (
	Mongo           DBType = "mongo"
	Psql            DBType = "psql"
	Device          DBType = "device"
	ConfigPath      string = "./.conf/todo-list-config.json"
	ConfigDirectory string = "./.conf/"
	StoragePath     string = "todo-list.json"
)

func (db DBType) IsValid() error {
	switch db {
	case Mongo, Psql:
		return nil
	}
	return errors.New("invalid DBType")
}

func (db DBType) String() string {
	return string(db)
}

type AppConfig struct {
	Db             DBType `json:"db"`
	DefaultStorage DBType `json:"defaultStorage"`
	Path           string `json:"path"`
}
