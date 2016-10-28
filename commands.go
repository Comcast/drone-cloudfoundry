package main

// API to execute against
func api(api API) []string {
	uri := api.URI
	require("api", uri)
	return []string{"api", uri}
}

// Login to cloud foundry
func login(credentials Credentials) []string {
	user, pass := credentials.User, credentials.Password
	require("user", user)
	require("password", pass)

	return []string{"auth", user, pass}
}

// Target an org/space
func target(vargs Target) []string {
	org, space := vargs.Org, vargs.Space
	require("org", org)
	require("space", space)
	return []string{"target", "-o", org, "-s", space}
}

// Push the application
func push(app App, route Route, flags Flags) []string {
	return combine(
		[]string{"push"},
		app.args(),
		route.args(),
		flags.args(),
	)
}
