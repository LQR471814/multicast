package store

import (
	"encoding/json"
	"fmt"
	"os"
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
	path, err := os.UserConfigDir()
	if err != nil {
		return err
	}

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
	path, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	data, err := json.Marshal(*store)
	if err != nil {
		return err
	}

	err = os.WriteFile(
		fmt.Sprintf("%v\\%v", path, StoreFilename),
		data,
		0644,
	)

	return err
}
