package checks

import (
	"os/exec"
	"strconv"
	"strings"
)

func win_firewall_check(ctx RuleContext) (bool, error) {
	err := exec.Command(
		"powershell.exe",
		"Get-NetFirewallRule",
		"-DisplayName \"FTX\"",
	).Run()

	return err == nil, err
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

	return true, err
}
