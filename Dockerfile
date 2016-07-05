# Docker image for the Drone Cloudfoundry plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-cloudfoundry
#     make deps build docker

FROM alpine:3.4

RUN echo "@testing http://dl-cdn.alpinelinux.org/alpine/edge/testing" | tee -a /etc/apk/repositories && \
  apk -U add \
    ca-certificates \
    git \
    cloudfoundry-cli@testing && \
  rm -rf /var/cache/apk/*

ADD drone-cloudfoundry /bin/
ENTRYPOINT ["/bin/drone-cloudfoundry"]
