FROM golang:1.13 AS builder

# create a working directory
RUN mkdir -p  /go/src/Kaplan-Go

COPY   .  /go/src/Kaplan-Go

WORKDIR /go/src/Kaplan-Go


ENV GO111MODULE=on

FROM alpine
RUN apk add --no-cache openssh

COPY --from=builder /go/src/Kaplan-Go/Kaplan-Go  /
COPY --from=builder /go/src/Kaplan-Go/configuration/config.json /


# Exposing port
EXPOSE 8085
# Run the widget-server  binary.
ENTRYPOINT ["/Kaplan-Go"]