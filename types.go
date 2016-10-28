package main

// Target of cf push
type Target struct {
	Org   string
	Space string
}

// Credentials of cf user
type Credentials struct {
	User     string
	Password string
}

// API target
type API struct {
	URI string
}

// App information
type App struct {
	Name     string
	Manifest string
	Path     string

	Command   string
	Buildpack string
	Disk      string
	Memory    string
	Instances int
}

// Route information
type Route struct {
	Hostname    string
	RandomRoute bool
	Domain      string
	NoRoute     bool
}

// Flags toggle true/false
type Flags struct {
	SkipSSL    bool
	NoStart    bool
	NoHostname bool
	NoManifest bool
}

// CloudFoundry plugin arguments
type CloudFoundry struct {
	API
	Target
	Credentials
	Flags
	Route
	App
}
