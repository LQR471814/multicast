package checks

import (
	"os/exec"
	"strconv"
	"strings"
)

func win_command_failed(err error) bool {
	if err != nil && err.Error() == "exit status 1" {
		return true
	}

	return false
}

func win_firewall_check(ctx RuleContext) (bool, error) {
	err := exec.Command(
		"powershell.exe",
		"Get-NetFirewallRule",
		"-DisplayName \"FTX\"",
	).Run()

	if win_command_failed(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func win_routing_check(ctx RuleContext) (bool, error) {
	out, err := exec.Command(
		"powershell.exe",
		"Get-NetRoute",
		"-InterfaceIndex",
		strconv.FormatInt(
			ctx.Interface,
			10,
		),
	).Output()

	if win_command_failed(err) {
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
