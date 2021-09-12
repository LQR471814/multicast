package multicast

import (
	"log"
	"net"

	"github.com/LQR471814/multicast/action"
	"github.com/LQR471814/multicast/common"
	"github.com/LQR471814/multicast/operations"
	"github.com/LQR471814/multicast/store"
)

func init() {
	err := store.LoadStore()
	if err != nil {
		log.Fatal(err)
	}
}

func Ping(group *net.UDPAddr, buf []byte) error {
	if store.Current() == store.DefaultStore() {
		return common.SetupRequired{}
	}

	err := operations.Ping(group, buf)
	return err
}

func Listen(group *net.UDPAddr, handler func(operations.MulticastPacket)) error {
	if store.Current() == store.DefaultStore() {
		return common.SetupRequired{}
	}

	err := operations.Listen(group, int(store.Current().Interface), handler)
	return err
}

func IsAdmin() bool {
	return action.IsAdmin()
}

func Setup(exec string, intf int) error {
	return action.Setup(exec, intf)
}

func Reset() error {
	return action.Reset(int(store.Current().Interface))
}
