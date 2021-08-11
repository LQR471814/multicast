package multicast

import (
	"log"
	"runtime"

	"multicast/checks"
	"multicast/reset"
	"multicast/setup"
)

var store *Store

func init() {
	store = &Store{}

	err := LoadStore(store)
	if err != nil {
		log.Fatal(err)
	}
}

func Check() (bool, error) {
	var result bool
	var err error

	ctx := checks.RuleContext{
		Interface: store.Interface,
	}

	switch runtime.GOOS {
	case "windows":
		result, err = checks.Check_Win(ctx)
	}

	return result, err
}

func Setup(intf int) error {
	switch runtime.GOOS {
	case "windows":
		return setup.Win32(intf)
	}

	return nil
}

func Reset(intf int) error {
	switch runtime.GOOS {
	case "windows":
		return reset.Win32(intf)
	}

	return nil
}
