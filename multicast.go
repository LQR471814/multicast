package multicast

import (
	"log"
	"net"
	"runtime"

	"github.com/LQR471814/multicast/checks"
	"github.com/LQR471814/multicast/common"
	"github.com/LQR471814/multicast/operations"
	"github.com/LQR471814/multicast/reset"
	"github.com/LQR471814/multicast/setup"
	"github.com/LQR471814/multicast/store"
)

func init() {
	err := store.LoadStore()
	if err != nil {
		log.Fatal(err)
	}
}

func Ping(group *net.UDPAddr, buf []byte) error {
	pingable, err := Check()
	if err != nil {
		return err
	}

	if !pingable {
		return common.SetupRequired{}
	}

	err = operations.Ping(group, buf)
	return err
}

func Listen(group *net.UDPAddr, handler func(operations.MulticastPacket)) error {
	listenable, err := Check()
	if err != nil {
		return err
	}

	if !listenable {
		return common.SetupRequired{}
	}

	err = operations.Listen(group, int(store.Current().Interface), handler)
	return err
}

func Check() (bool, error) { //? Returns false if setup is required
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

func Setup(exec string, intf int) error {
	switch runtime.GOOS {
	case "windows":
		return setup.Win(exec, intf)
	}

	return nil
}

func Reset() error {
	switch runtime.GOOS {
	case "windows":
		return reset.Win(int(store.Current().Interface))
	}

	return nil
}
