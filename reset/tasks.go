package reset

import (
	"os/exec"

	"github.com/LQR471814/multicast/common"
	"github.com/LQR471814/multicast/setup"
)

func Win_routing_reset(intf int) ([]byte, error) {
	return setup.Win_routing_cfg(intf, 256)
}

func Win_firewall_reset() ([]byte, error) {
	if !common.Win32_IsAdmin() {
		return []byte{}, common.MissingPrivileges{}
	}

	out, err := exec.Command(
		"powershell",
		"Remove-NetFirewallRule",
		"-DisplayName", common.FIREWALL_NAME,
	).Output()

	return out, err
}
