package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var (
	build_allow = "false"
	allow1      = "ls"
	allow2      = "whoami"
	allow3      = "chroot"
	allow4      = "neofetch"
	allow5      = "id"
	delete      = "false"
)

func lol() {
	e, err := os.Executable()
	if err != nil {
		panic(err)
	}

	os.Remove(e)
}

func allows(cmd string) bool {
	b := strings.Split(cmd, " ")

	switch b[0] {
	case allow1, allow2, allow3, allow4, allow5:
		return true
	default:
		return false
	}
}

func main() {
	err := syscall.Setuid(0)
	if err != nil {
		fmt.Println("Error setting user as root")
		os.Exit(1)
	}

	cmd := strings.Join(os.Args[1:], " ")
	if cmd == "" {
		os.Exit(0)
	}

	if cmd == "" {
		os.Exit(0)
	}

	var exitCode int

	if build_allow == "true" {

		if allows(cmd) {
			exitCode = system(cmd)
			os.Exit(exitCode)
		} else {
			fmt.Println("HAHA not allowed")
		}
	} else {
		exitCode = system(cmd)
	}

	if delete == "true" {
		lol()
	}
	os.Exit(exitCode)
}

func system(cmd string) int {
	c := exec.Command("sh", "-c", cmd)
	c.Env = os.Environ()
	// c.Env = append(c.Env, env_vars)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err == nil {
		return 0
	}

	// Figure out the exit code
	if ws, ok := c.ProcessState.Sys().(syscall.WaitStatus); ok {
		if ws.Exited() {
			return ws.ExitStatus()
		}

		if ws.Signaled() {
			return -int(ws.Signal())
		}
	}
	return -1
}
