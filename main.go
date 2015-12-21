package main

import (
	"fmt"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone Cloud Foundry Plugin built at %s\n", buildDate)

	cfargs := CloudFoundry{}
	workspace := drone.Workspace{}
	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &cfargs)
	plugin.MustParse()

	run(api(cfargs.API))

	run(login(cfargs.Credentials))

	run(target(cfargs.Target))

	run(deploy(workspace, cfargs.App, cfargs.Route, cfargs.Flags))
}
