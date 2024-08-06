package configs

import "errors"

type DBType string

const (
	Mongo           DBType = "mongo"
	Psql            DBType = "psql"
	Device          DBType = "device"
	ConfigPath      string = "./.conf/todo-list-config.json"
	ConfigDirectory string = "./.conf/"
	StoragePath     string = "todo-list.json"
)

var DefaultConfig AppConfig = AppConfig{
	Db:             Device,
	DefaultStorage: Device,
	Path:           StoragePath,
}

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
