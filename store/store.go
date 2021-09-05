package store

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kirsle/configdir"
)

const (
	StoreFilename = "go-multicast-config.json"
)

type Store struct {
	Interface int64
}

var store = &Store{}

func Current() Store {
	return *store
}

func DefaultStore() Store {
	return Store{
		Interface: -1,
	}
}

func UpdateStore(newStore Store) {
	*store = newStore
	WriteStore()
}

func LoadStore() error {
	path := configdir.SystemConfig()[0]

	data, err := os.ReadFile(fmt.Sprintf("%v\\%v", path, StoreFilename))
	if os.IsNotExist(err) {
		UpdateStore(DefaultStore())
	} else if err != nil {
		return err
	} else {
		err = json.Unmarshal(data, store)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteStore() error {
	path := configdir.SystemConfig()[0]

	data, err := json.Marshal(*store)
	if err != nil {
		return err
	}

	configDir := fmt.Sprintf("%v\\%v", path, StoreFilename)

	f, err := os.Create(configDir)
	if err != nil {
		return err
	}

	defer f.Close()

	f.Write(data)
	f.Sync()

	return err
}
