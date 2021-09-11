package setup

import (
	"github.com/LQR471814/multicast/common"
	"github.com/LQR471814/multicast/store"
)

func Win(exec string, intf int) error {
	if !common.Win_IsAdmin() {
		return common.MissingPrivileges{}
	}

	_, err := Win_routing_cfg(intf, 1)
	if err != nil {
		return err
	}

	_, err = Win_firewall_setup(exec)
	if err != nil {
		return err
	}

	store.UpdateStore(store.Store{
		Interface: int64(intf),
	})

	return nil
}
