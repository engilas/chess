//go:build unix
// +build unix

package uci

import "syscall"

func getAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setpgid: true,
	}
}
