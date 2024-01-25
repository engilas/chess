//go:build !windows
// +build !windows

package uci

import "syscall"

func getAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setpgid: true,
	}
}
