package main

import "strconv"

func param(arg string, value string) []string {
	if value == "" {
		return nil
	}
	return []string{arg, value}
}
func arg(arg string) []string {
	if arg == "" {
		return nil
	}
	return []string{arg}
}
func flag(arg string, value bool) []string {
	if value {
		return []string{arg}
	}
	return nil
}

func (app *App) args() []string {
	return combine(
		arg(app.Name),
		param("-f", app.Manifest),
		param("-p", app.Path),
		param("-i", func(instances int) string {
			if instances > 0 {
				return strconv.Itoa(instances)
			}
			return ""
		}(app.Instances)),
		param("-k", app.Disk),
		param("-b", app.Buildpack),
		param("-c", app.Command),
		param("-m", app.Memory),
	)
}

func (flags *Flags) args() []string {
	return combine(
		flag("--no-start", flags.NoStart),
		flag("--no-hostname", flags.NoHostname),
		flag("--no-manifest", flags.NoManifest),
		flag("--skip-ssl-validation", flags.SkipSSL),
	)
}

func (route *Route) args() []string {
	return combine(
		param("-d", route.Domain),
		param("-n", route.Hostname),
		flag("--no-route", route.NoRoute),
		flag("--random-route", route.RandomRoute),
	)
}
