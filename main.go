package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

func main() {
	cfargs := CloudFoundry{}
	workspace := drone.Workspace{}
	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &cfargs)
	plugin.Parse()

	run(api(cfargs.API))

	run(login(cfargs.Credentials))

	run(target(cfargs.Target))

	run(deploy(workspace, cfargs.App, cfargs.Route, cfargs.Flags))
}

// checks that a field has been set
func nnull(field string, val string) {
	if val == "" {
		fmt.Fprint(os.Stderr, "`", field, "` is a required field\n")
		os.Exit(1)
	}
}

// cf api
func api(vargs API) *exec.Cmd {
	uri := vargs.URI
	nnull("api", uri)
	fmt.Printf("Target api %s\n", uri)
	return exec.Command("cf", "api", uri)
}

// cf login
func login(vargs Credentials) *exec.Cmd {
	user, pass := vargs.User, vargs.Password
	nnull("user", user)
	nnull("password", pass)

	fmt.Println("Logging in...")
	return exec.Command("cf", "auth", user, pass)
}

// cf target
func target(vargs Target) *exec.Cmd {
	org, space := vargs.Org, vargs.Space
	nnull("org", org)
	nnull("space", space)
	fmt.Printf("Targeting %s:%s...\n", org, space)
	return exec.Command("cf", "target", "-o", org, "-s", space)
}

// cf deploy
func deploy(workspace drone.Workspace, app App, route Route, flags Flags) *exec.Cmd {
	fmt.Println("Deploy")
	args := combine(
		[]string{"push"},
		parseApp(app),
		parseRoute(route),
		parseFlags(flags),
	)

	cmd := exec.Command("cf", args...)
	cmd.Dir = workspace.Path
	return cmd
}
func parseApp(app App) []string {
	args := []string{}
	if app.Name != "" {
		return append(args, app.Name)
	}
	return args
}
func parseFlags(flags Flags) []string {
	args := []string{}
	if flags.NoStart {
		args = append(args, "--no-start")
	}
	if flags.NoHostname {
		args = append(args, "--no-hostname")
	}
	if flags.NoManifest {
		args = append(args, "--no-manifest")
	}
	return args
}
func parseRoute(route Route) []string {
	args := []string{}
	if route.Domain != "" {
		args = append(args, "-d", route.Domain)
	}
	if route.Hostname != "" {
		args = append(args, "-n", route.Hostname)
	}
	if route.NoRoute {
		args = append(args, "--no-route")
	}
	if route.RandomRoute {
		args = append(args, "--random-route")
	}
	return args
}
func combine(strs ...[]string) []string {
	out := []string{}
	for _, slice := range strs {
		out = append(out, slice...)
	}
	return out
}

// run a shell command
func run(cmd *exec.Cmd) {
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
