# Docker image for the Drone Cloud Foundry plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-cloudfoundry
#     make deps build docker

FROM golang:1.5

RUN go get github.com/cloudfoundry/cli/main
RUN go build -o /bin/cf github.com/cloudfoundry/cli/main
ADD drone-cloudfoundry /bin/
ENTRYPOINT ["/bin/drone-cloudfoundry"]
