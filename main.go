package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func lol() {
	e, err := os.Executable()
	if err != nil {
		panic(err)
	}

	os.Remove(e)
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

	exitCode := system(cmd)
	lol()
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
