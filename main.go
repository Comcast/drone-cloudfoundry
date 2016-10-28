package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	version string
)

func main() {
	app := cli.NewApp()
	app.Name = "drone-cloudfoundry"
	app.Usage = "drone-cloudfoundry usage"
	app.Action = run
	app.Version = version

	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:   "api.uri",
			Usage:  "api uri",
			EnvVar: "PLUGIN_API",
		},
		cli.StringFlag{
			Name:   "credentials.user",
			Usage:  "cf username",
			EnvVar: "CF_USER,PLUGIN_USER",
		},
		cli.StringFlag{
			Name:   "credentials.password",
			Usage:  "cf password",
			EnvVar: "CF_PASSWORD,PLUGIN_PASSWORD",
		},
		cli.StringFlag{
			Name:   "target.org",
			Usage:  "target org",
			EnvVar: "CF_ORG,PLUGIN_ORG",
		},
		cli.StringFlag{
			Name:   "target.space",
			Usage:  "target space",
			EnvVar: "CF_SPACE,PLUGIN_SPACE",
		},
		cli.StringFlag{
			Name:   "app.name",
			Usage:  "cf app name",
			EnvVar: "PLUGIN_NAME",
		},
		cli.StringFlag{
			Name:   "app.manifest",
			Usage:  "path to manifest yaml",
			EnvVar: "PLUGIN_MANIFEST",
		},
		cli.StringFlag{
			Name:   "app.path",
			Usage:  "app path",
			EnvVar: "PLUGIN_PATH",
		},
		cli.StringFlag{
			Name:   "app.command",
			Usage:  "startup command",
			EnvVar: "PLUGIN_COMMAND",
		},
		cli.StringFlag{
			Name:   "app.buildpack",
			Usage:  "custom buildpack",
			EnvVar: "PLUGIN_BUILDPACK",
		},
		cli.StringFlag{
			Name:   "app.disk",
			Usage:  "disk limit",
			EnvVar: "PLUGIN_DISK",
		},
		cli.StringFlag{
			Name:   "app.memory",
			Usage:  "memory limit",
			EnvVar: "PLUGIN_MEMORY",
		},
		cli.IntFlag{
			Name:   "app.instances",
			Usage:  "number of instances",
			EnvVar: "PLUGIN_INSTANCES",
		},
		cli.StringFlag{
			Name:   "route.hostname",
			Usage:  "hostname",
			EnvVar: "PLUGIN_HOSTNAME",
		},
		cli.BoolFlag{
			Name:   "route.random",
			Usage:  "create a random route",
			EnvVar: "PLUGIN_RANDOM_ROUTE",
		},
		cli.StringFlag{
			Name:   "route.domain",
			Usage:  "domain",
			EnvVar: "PLUGIN_DOMAIN",
		},
		cli.BoolFlag{
			Name:   "route.noroute",
			Usage:  "do not map a route to this app and remove previous routes",
			EnvVar: "PLUGIN_NO_ROUTE",
		},
		cli.BoolFlag{
			Name:   "flags.skipssl",
			Usage:  "do not map a route to this app and remove previous routes",
			EnvVar: "PLUGIN_SKIP_SSL",
		},
		cli.BoolFlag{
			Name:   "flags.nostart",
			Usage:  "do not start an app after pushing",
			EnvVar: "PLUGIN_NO_START",
		},
		cli.BoolFlag{
			Name:   "flags.nohostname",
			Usage:  "map the root domain to this app",
			EnvVar: "PLUGIN_NO_HOSTNAME",
		},
		cli.BoolFlag{
			Name:   "flags.nomanifest",
			Usage:  "ignore manifest file",
			EnvVar: "PLUGIN_NO_MANIFEST",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) {

	plugin := Plugin{
		Target: Target{
			Org:   c.String("target.org"),
			Space: c.String("target.space"),
		},
		Credentials: Credentials{
			User:     c.String("credentials.user"),
			Password: c.String("credentials.password"),
		},
		API: API{
			URI: c.String("api.uri"),
		},
		App: App{
			Name:     c.String("app.name"),
			Manifest: c.String("app.manifest"),
			Path:     c.String("app.path"),

			Command:   c.String("app.command"),
			Buildpack: c.String("app.buildpack"),
			Disk:      c.String("app.disk"),
			Memory:    c.String("app.memory"),
			Instances: c.Int("app.instances"),
		},
		Route: Route{
			Hostname:    c.String("route.hostname"),
			RandomRoute: c.Bool("route.random"),
			Domain:      c.String("route.domain"),
			NoRoute:     c.Bool("route.noroute"),
		},
		Flags: Flags{
			SkipSSL:    c.Bool("flags.skipssl"),
			NoStart:    c.Bool("flags.nostart"),
			NoHostname: c.Bool("flags.nohostname"),
			NoManifest: c.Bool("flags.nomanifest"),
		},
	}

	fmt.Println(plugin)

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
