# drone-cloudfoundry

[Drone](https://github.com/drone/drone) plugin to deploy or update a project on [Cloud Foundry](https://www.cloudfoundry.org/).
For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

### Example

```sh
docker run --rm \
    -e PLUGIN_API=<api> \
    -e PLUGIN_USER=<username> \
    -e PLUGIN_PASSWORD=<password> \
    -e PLUGIN_ORG=<org> \
    -e PLUGIN_SPACE=<space> \
    -v $(pwd):$(pwd) \
    -w $(pwd) \
    plugins/drone-cloudfoundry
```

## Docker

Build the docker image with the following commands:

```sh
docker build --rm=true -t cheslip/drone-cloudfoundry .
```
