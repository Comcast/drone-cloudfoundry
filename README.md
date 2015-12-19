# drone-cloudfoundry

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-cloudfoundry/status.svg)](http://beta.drone.io/drone-plugins/drone-cloudfoundry)
[![](https://badge.imagelayers.io/plugins/drone-cloudfoundry:latest.svg)](https://imagelayers.io/?images=plugins/drone-cloudfoundry:latest 'Get your own badge on imagelayers.io')

Drone plugin for deploying to Cloud Foundry

See [DOCS.md](DOCS.md) for `.drone.yml` usage

## Usage

```
./drone-cloudfoundry <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
      "api": "api.run.pivotal.io",
      "org": "my-org",
      "space": "dev",
      "user": "johndoe",
      "password": "pa55word",
      "name": "test-cf-deploy",
      "manifest": "manifest.yml",
      "path": ".",
      "command": "npm start",
      "buildpack": "nodejs",
      "disk": "128",
      "memory": "64",
      "instances": 1,
      "hostname": "",
      "random-route": false,
      "domain": "apps.pivotal.io",
      "no-route": false,
      "skip-ssl-validation": false,
      "no-start": false,
      "no-hostname": false,
      "no-manifest": false
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```
make deps build docker
```

### Example

```sh
docker run -i plugins/drone-cloudfoundry <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
    }
}
EOF
```
