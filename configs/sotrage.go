package configs

import "errors"

type DBType string

const (
	Mongo       DBType = "mongo"
	Psql        DBType = "psql"
	Device      DBType = "device"
	ConfigPath  string = "todo-list-config.json"
	StoragePath string = "todo-list.json"
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
