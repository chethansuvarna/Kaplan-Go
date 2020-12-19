# Stage 1
FROM golang:1.13 AS builder
# Creating working directory
RUN mkdir -p  /go/src/Kaplan-Go

# Copying source code to repository
COPY   .  /go/src/Kaplan-Go
WORKDIR /go/src/Kaplan-Go

# Installing ca certificates
RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates && rm -rf /var/lib/apt/lists/*
ENV GO111MODULE=on

# Creating go binary
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Kaplan-Go
# Stage 2
FROM alpine
RUN apk add --no-cache openssh

# Copy ca certificates from builder
COPY --from=builder  /etc/ssl/certs /etc/ssl/certs

# Copy our static executable and dependencies from builder
COPY --from=builder /go/src/Kaplan-Go  /
COPY --from=builder /go/src/Kaplan-Go/configuration/config.json  /

# Exposing port
EXPOSE 8080

# Run the Kaplan-Go  binary.
ENTRYPOINT ["/Kaplan-Go"]