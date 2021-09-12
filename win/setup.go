package win

import (
	"os/exec"
	"strconv"

	"github.com/LQR471814/multicast/common"
	"github.com/LQR471814/multicast/store"
)

func routing_setup(intf, metric int) ([]byte, error) {
	out, err := exec.Command( //? Set route
		"netsh", "interface",
		"ipv4", "set", "route",
		"224.0.0.0/4",
		"interface="+strconv.Itoa(intf),
		"siteprefixlength=0",
		"metric=1", "publish=yes",
		"store=persistent",
	).Output()

	return out, err
}

func firewall_setup(path string) ([]byte, error) {
	out, err := exec.Command(
		"netsh", "advfirewall",
		"firewall", "add",
		"rule", "name="+common.FIREWALL_NAME,
		"program="+path,
		"protocol=udp", "dir=in",
		"enable=yes", "action=allow",
		"profile=Any",
	).Output()

	return out, err
}

func Setup(exec string, intf int) error {
	if !IsAdmin() {
		return common.MissingPrivileges{}
	}

	_, err := routing_setup(intf, 1)
	if err != nil {
		return err
	}

	_, err = firewall_setup(exec)
	if err != nil {
		return err
	}

	store.UpdateStore(store.Store{
		Interface: int64(intf),
	})

	return nil
}
