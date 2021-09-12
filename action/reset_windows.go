package action

import (
	"os/exec"

	"github.com/LQR471814/multicast/common"
)

func win_routing_reset(intf int) ([]byte, error) {
	return routing_setup(intf, 256)
}

func win_firewall_reset() ([]byte, error) {
	if !IsAdmin() {
		return []byte{}, common.MissingPrivileges{}
	}

	out, err := exec.Command(
		"powershell",
		"Remove-NetFirewallRule",
		"-DisplayName", common.FIREWALL_NAME,
	).Output()

	return out, err
}

func Reset(intf int) error {
	_, err := win_routing_reset(intf)
	if err != nil {
		return err
	}

	_, err = win_firewall_reset()

	return err
}
