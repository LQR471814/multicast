package multicast

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

func DefaultStore() Store {
	return Store{
		Interface: -1,
	}
}

func LoadStore(store *Store) error {
	path, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(fmt.Sprintf("%v\\%v", path, StoreFilename))
	if err != nil {
		*store = DefaultStore()
		return err
	}

	err = json.Unmarshal(data, store)
	return err
}

func WriteStore(store *Store) error {
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
