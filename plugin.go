package main

import (
	"io"
	"os"
	"os/exec"
)

type (
	Plugin struct {
		Target      Target
		Credentials Credentials
		API         API
		App         App
		Route       Route
		Flags       Flags
	}
)

func (p Plugin) Exec() error {
	cli := cfcli{}
	cli.Exec(
		api(p.API)...)
	cli.Exec(
		login(p.Credentials)...)
	cli.Exec(
		target(p.Target)...)
	cli.Exec(
		push(p.App, p.Route, p.Flags)...)
	return nil
}

type cfcli struct {
	Dir string
}

func (cli cfcli) Command(args ...string) *exec.Cmd {
	cmd := exec.Command("cf", args...)
	cmd.Dir = cli.Dir
	return cmd
}

func (cli cfcli) Exec(args ...string) {
	cmd := cli.Command(args...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	cmd.Start()
	io.Copy(os.Stdout, stdout)
	io.Copy(os.Stderr, stderr)
	err := cmd.Wait()

	if err != nil {
		os.Exit(1)
	}
}
