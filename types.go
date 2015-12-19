package main

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

// App information
type App struct {
	Name     string `json:"name"`
	Manifest string `json:"manifest"`
	Path     string `json:"path"`

	Command   string `json:"command"`
	Buildpack string `json:"buildpack"`
	Disk      string `json:"disk"`
	Memory    string `json:"memory"`
	Instances int    `json:"instances"`
}

// Route information
type Route struct {
	Hostname    string `json:"hostname"`
	RandomRoute bool   `json:"random-route"`
	Domain      string `json:"domain"`
	NoRoute     bool   `json:"no-route"`
}

// Flags toggle true/false
type Flags struct {
	SkipSSL    bool `json:"skip-ssl-validation"`
	NoStart    bool `json:"no-start"`
	NoHostname bool `json:"no-hostname"`
	NoManifest bool `json:"no-manifest"`
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
