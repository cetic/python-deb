from golang:1.13.5-alpine as build
run apk add git
add setup2control.go setup2control.go
run go build setup2control.go
cmd bash

from ubuntu:bionic
RUN apt-get update -y && apt-get install -y software-properties-common python3-pip python3-venv python3-setuptools curl
RUN add-apt-repository ppa:rmescandon/yq
RUN apt-get -y update && apt-get install -y yq
COPY --from=build /go/setup2control /usr/bin
cmd bash
