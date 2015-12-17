# Docker image for the Drone Cloud Foundry plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-cloudfoundry
#     make deps build docker

FROM alpine:3.2

RUN apk update && apk add ca-certificates
RUN apk add go git

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GOPATH=/go

RUN mkdir /go
RUN go get github.com/cloudfoundry/cli/main
RUN go build -o /bin/cf github.com/cloudfoundry/cli/main

ADD drone-cloudfoundry /bin/
ENTRYPOINT ["/bin/drone-cloudfoundry"]
