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
	}

	lines := []string{}
	for _, line := range strings.Split(string(out), "\r\n") {
		if strings.Contains(line, "224.0.0.0") {
			lines = append(lines, line)
		}
	}

	for i := 0; i < len(lines); i++ {
		if strings.Fields(lines[i])[3] != "256" {
			return false, err
		}
	}

	return true, nil
}
