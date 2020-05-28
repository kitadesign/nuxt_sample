FROM golang:1.12.4-alpine3.9 as build-step

RUN apk add --update --no-cache ca-certificates git

WORKDIR /usr/local/backend/src
RUN go get -u github.com/VojtechVitek/rerun/cmd/rerun
