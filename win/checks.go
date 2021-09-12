//go:build windows
// +build windows

package win

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/LQR471814/multicast/common"
)

func command_failed(err error) bool {
	if err != nil && err.Error() == "exit status 1" {
		return true
	}

	return false
}

func firewall_check(ctx common.RuleContext) (bool, error) {
	err := exec.Command(
		"powershell.exe",
		"Get-NetFirewallRule",
		"-DisplayName \"FTX\"",
	).Run()

	if command_failed(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func routing_check(ctx common.RuleContext) (bool, error) {
	out, err := exec.Command(
		"powershell.exe",
		"Get-NetRoute",
		"-InterfaceIndex",
		strconv.FormatInt(
			ctx.Interface,
			10,
		),
	).Output()

	if command_failed(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	lines := []string{}
	for _, ln := range strings.Split(string(out), "\r\n") {
		if strings.Contains(ln, "224.0.0.0") {
			lines = append(lines, ln)
		}
	}

	for _, ln := range lines {
		if strings.Fields(ln)[3] != lines[0] {
			return true, nil
		}
	}

	return false, nil
}

func Check(ctx common.RuleContext) (bool, error) {
	return common.All(ctx, []common.Rule{ //? Returns true if all rules pass
		firewall_check,
		routing_check,
	})
}
