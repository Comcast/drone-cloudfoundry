# Docker image for the Drone Cloud Foundry plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-cloudfoundry
#     make deps build docker

FROM alpine:3.2

RUN apk update && \
  apk add ca-certificates && \
  rm -rf /var/cache/apk/*

ADD cf /bin/
ADD drone-cloudfoundry /bin/
ENTRYPOINT ["/bin/drone-cloudfoundry"]
