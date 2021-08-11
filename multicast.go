package multicast

import (
	"log"
	"runtime"

	"github.com/LQR471814/multicast/checks"
	"github.com/LQR471814/multicast/common"
	"github.com/LQR471814/multicast/operations"
	"github.com/LQR471814/multicast/reset"
	"github.com/LQR471814/multicast/setup"
	"github.com/LQR471814/multicast/store"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := store.LoadStore()
	if err != nil {
		log.Fatal(err)
	}
}

func Ping(buf []byte) error {
	pingable, err := Check()
	if err != nil {
		return err
	}

	if !pingable {
		return common.SetupRequired{}
	}

	err = operations.Ping(buf)
	return err
}

func Listen(handler func(operations.MulticastPacket)) error {
	listenable, err := Check()
	if err != nil {
		return err
	}

	if !listenable {
		return common.SetupRequired{}
	}

	err = operations.Listen(int(store.Current().Interface), handler)
	return err
}

func Check() (bool, error) {
	var result bool
	var err error

	ctx := checks.RuleContext{
		Interface: store.Current().Interface,
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
		return setup.Win(intf)
	}

	return nil
}

func Reset(intf int) error {
	switch runtime.GOOS {
	case "windows":
		return reset.Win(intf)
	}

	return nil
}
