FROM alpine
MAINTAINER Cory Heslip <cheslip@mac.com>

RUN apk -Uuv add ca-certificates openssl bash

ENV CF_VERSION 6.19.0
RUN wget -qO - "https://cli.run.pivotal.io/stable?release=linux64-binary&version=${CF_VERSION}" | tar -xz -C /bin/

ADD deploy.sh /bin/
RUN chmod +x /bin/deploy.sh

ENTRYPOINT /bin/deploy.sh
