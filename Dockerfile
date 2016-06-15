# Docker image for the Drone Cloud Foundry plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-cloudfoundry
#     make deps build docker

FROM alpine:3.2

RUN apk update && \
  apk add ca-certificates && \
  rm -rf /var/cache/apk/*

ENV CF_VERSION 6.19.0
RUN wget -qO - "https://cli.run.pivotal.io/stable?release=linux64-binary&version=${CF_VERSION}" | tar -xz -C /bin/

ADD drone-cloudfoundry /bin/
ENTRYPOINT ["/bin/drone-cloudfoundry"]
