# drone-cloudfoundry

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-cloudfoundry/status.svg)](http://beta.drone.io/drone-plugins/drone-cloudfoundry)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-cloudfoundry/coverage.svg)](https://aircover.co/drone-plugins/drone-cloudfoundry)
[![](https://badge.imagelayers.io/plugins/drone-cloudfoundry:latest.svg)](https://imagelayers.io/?images=plugins/drone-cloudfoundry:latest 'Get your own badge on imagelayers.io')

Drone plugin to deploy or update a project on Cloud Foundry. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

### Example

```sh
docker run --rm \
    -e PLUGIN_API=<api>
    -e PLUGIN_USER=<username>
    -e PLUGIN_PASSWORD=<password>
    -e PLUGIN_ORG=<org>
    -e PLUGIN_SPACE=<space>
    -v $(pwd):$(pwd) \
    -w $(pwd)
    plugins/drone-cloudfoundry
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t plugins/drone-cloudfoundry .
```
