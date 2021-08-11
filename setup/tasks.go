package setup

import (
	"os/exec"
	"strconv"

	"multicast/common"
)

func Win_routing_cfg(intf, metric int) ([]byte, error) {
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

func Win_firewall_setup(path string) ([]byte, error) {
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
