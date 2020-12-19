FROM golang:1.13 AS builder

# create a working directory
RUN mkdir -p  /go/src/Kaplan

COPY   .  /go/src/Kaplan-Go

WORKDIR /go/src/Kaplan/Kaplan-Go

# RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates && rm -rf /var/lib/apt/lists/*

ENV GO111MODULE=on
RUN go mod init && go clean

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Kaplan-Go

FROM alpine
RUN apk add --no-cache openssh

COPY --from=builder /go/src/Kaplan-Go/Kaplan-Go  /
COPY --from=builder /go/src/Kaplan-Go/configuration/config.json /


# Exposing port
EXPOSE 8085
# Run the widget-server  binary.
ENTRYPOINT ["/Kaplan-Go"]