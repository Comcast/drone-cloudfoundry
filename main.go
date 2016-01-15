package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone Cloud Foundry Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	cfargs := CloudFoundry{}

	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &cfargs)
	plugin.MustParse()

	cli := cfcli{
		Dir: workspace.Path,
	}

	cli.Exec(
		api(cfargs.API)...)
	cli.Exec(
		login(cfargs.Credentials)...)
	cli.Exec(
		target(cfargs.Target)...)
	cli.Exec(
		push(
			workspace,
			cfargs.App,
			cfargs.Route,
			cfargs.Flags)...)
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
