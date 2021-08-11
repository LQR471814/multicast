package setup

import (
	"log"
	"os"

	"github.com/LQR471814/multicast/common"
)

func Win32(intf int) error {
	if !common.Win32_IsAdmin() {
		return common.MissingPrivileges{}
	}

	out, err := Win_routing_cfg(intf, 1)
	if err != nil {
		return err
	}

	log.Default().Println(string(out))

	path, err := os.Executable()
	if err != nil {
		return err
	}

	out, err = Win_firewall_setup(path)
	if err != nil {
		return err
	}

	log.Default().Println(string(out))

	return nil
}

// func rerunElevated() {
// 	verb := "runas"
// 	exe, _ := os.Executable()
// 	cwd, _ := os.Getwd()
// 	args := strings.Join(os.Args[1:], " ")

// 	verbPtr, _ := syscall.UTF16PtrFromString(verb)
// 	exePtr, _ := syscall.UTF16PtrFromString(exe)
// 	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
// 	argPtr, _ := syscall.UTF16PtrFromString(args)

// 	var showCmd int32 = 1

// 	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
// 	if err != nil {
// 		log.Default().Println(err)
// 	}
// }
