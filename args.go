package main

func (app *App) args() []string {
	args := []string{}
	if app.Name != "" {
		return append(args, app.Name)
	}
	return args
}
func (flags *Flags) args() []string {
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
func (route *Route) args() []string {
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
