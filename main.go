package main

import (
  "io"
  "os"
	"fmt"
  "os/exec"
  "github.com/drone/drone-go/drone"
  "github.com/drone/drone-go/plugin"
  "strings"
)

type CloudFoundry struct {
  Api string `json:"api"`
  User string `json:"user"`
  Password string `json:"password"`
  Org string `json:"org"`
  Space string `json:"space"`
}

func main() {
  repo := drone.Repo{}
  build := drone.Build{}
  workspace := drone.Workspace{}
  vargs := CloudFoundry{}

  plugin.Param("repo", &repo)
  plugin.Param("build", &build)
  plugin.Param("workspace", &workspace)
  plugin.Param("vargs", &vargs)
  plugin.Parse()

  if vargs.Api == "" || vargs.User == "" || vargs.Password == "" {
    fmt.Fprintln(os.Stderr, "required")
    os.Exit(1)
  }
  var cmd *exec.Cmd

  // login
  cmd = login(vargs)
  run(*cmd)

  // cf push
  cmd = exec.Command("cf", "push")
  fmt.Println(strings.Join(cmd.Args, " "))
  run(*cmd)
}
func run(cmd exec.Cmd) error {
  stdout, _ := cmd.StdoutPipe()
  stderr, _ := cmd.StderrPipe()

  cmd.Start()

  go io.Copy(os.Stdout, stdout)
  go io.Copy(os.Stderr, stderr)
  return cmd.Wait()
}
func login(vargs CloudFoundry) *exec.Cmd {
  return exec.Command("cf", "login", "-a", vargs.Api, "-u", vargs.User, "-p", vargs.Password)
}
