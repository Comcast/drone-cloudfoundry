package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func combine(strs ...[]string) []string {
	out := []string{}
	for _, slice := range strs {
		out = append(out, slice...)
	}
	return out
}

// run a shell command
func runCmd(cmd *exec.Cmd) {
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

// Require that a field has been set
func require(field string, val interface{}) {
	if val == nil {
		reject(field)
	}
	if val == "" {
		reject(field)
	}
}
func reject(field string) {
	fmt.Fprint(os.Stderr, "`", field, "` is a required field\n")
	os.Exit(1)
}
