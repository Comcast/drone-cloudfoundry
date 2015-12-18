package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

// Target of cf push
type Target struct {
	Org   string `json:"org"`
	Space string `json:"space"`
}

// Credentials of cf user
type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// API target
type API struct {
	URI string `json:"api"`
}

// CloudFoundry plugin arguments
type CloudFoundry struct {
	API
	Target
	Credentials
}

func main() {
	cfargs := CloudFoundry{}
	workspace := drone.Workspace{}
	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &cfargs)
	plugin.Parse()

	run(api(cfargs.API))

	run(login(cfargs.Credentials))

	run(target(cfargs.Target))

	run(deploy(workspace))
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
func deploy(workspace drone.Workspace) *exec.Cmd {
	fmt.Println("Deploy")
	cmd := exec.Command("cf", "push", workspace.Path)
	cmd.Dir = workspace.Path
	return cmd
}

// run runs a shell command
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
