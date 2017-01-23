Use this plugin for deploying an application to Cloud Foundry. This is a bash script
that uses the Cloud Foundry CLI to login and deploy. More info and documentation about 
the CLI can be found [here](https://github.com/cloudfoundry/cli)

## Example

```yaml
deploy:
  image: cheslip/drone-cloudfoundry
  api: api.run.pivotal.io
  org: pivotal
  space: production
  when:
    branch: master
```

## Options

| Option | Description | Example |
| ------ | ----------- | ---- |
| `api` (**required**) | Target API | `api.run.pivotal.io` |
| `org` (**required**) | Target Org | `xyz-org` |
| `space` (**required**) | Target Space | `development` |
| `user` (**required**) | Auth username | `john@doe.com` |
| `password` (**required**) | Auth password | `supersecure` |
| `name` | Override application name | `app-canary` |
| `buildpack` | Custom buildpack | `https://....` |
| `command` | Startup command | `start-script.sh` |
| `domain` | Domain | `example.com` |
| `manifest` | Path to manifest | `test.manifest.yml` |
| `docker_image` | Image name | `cheslip/drone-cloudfoundry` |
| `instances` | Number of instances | `4` |
| `disk` | Disk limit | `256M` |
| `memory` | Memory limit | `1G` |
| `hostname` | Hostname | `my-subdomain` |
| `path` | App path | `build/assets` |
| `stack` | Stack to use | `cflinuxfs2` |
| `timeout` | App start timeout | `60` |
| `health_check_type` | Application health check type | `port` |
| `route_path` | Path for the route | `about` |
| `no_hostname` | Map the root domain to this app | `false` |
| `no_manifest` | Ignore manifest file | `false` |
| `no_route` | Do not map a route to this app and remove routes from previous pushes of this app. | `false` |
| `no_start` | Do not start an app after pushing | `false` |
| `random_route` | Create a random route for this app | `false` |

## Secrets

All required fields are also supported by secrets:

- `CF_API`
- `CF_USER`
- `CF_PASSWORD`
- `CF_ORG`
- `CF_SPACE`
