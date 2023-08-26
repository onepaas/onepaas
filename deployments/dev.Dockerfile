ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine

RUN apk update && \
	apk add \
            build-base \
            curl \
            git \
			make \
			openssh

WORKDIR /root/src

CMD tail -f /dev/null
