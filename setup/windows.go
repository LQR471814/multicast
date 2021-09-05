package setup

import (
	"os"

	"github.com/LQR471814/multicast/common"
	"github.com/LQR471814/multicast/store"
)

func Win(intf int) error {
	if !common.Win_IsAdmin() {
		return common.MissingPrivileges{}
	}

	_, err := Win_routing_cfg(intf, 1)
	if err != nil {
		return err
	}

	path, err := os.Executable()
	if err != nil {
		return err
	}

	_, err = Win_firewall_setup(path)
	if err != nil {
		return err
	}

	store.UpdateStore(store.Store{
		Interface: int64(intf),
	})

	return nil
}
