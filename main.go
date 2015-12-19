package main

import (
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
