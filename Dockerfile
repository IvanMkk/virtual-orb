FROM golang:alpine AS build-env

COPY . .
RUN GOPATH=/ go mod download
RUN GOPATH=/ GO111MODULE=on GOOS=linux GOARCH=amd64 \
    go build -o /bin/virtual_orb

FROM alpine:latest
COPY --from=build-env /bin/virtual_orb /virtual_orb
