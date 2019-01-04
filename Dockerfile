FROM golang:alpine
MAINTAINER nagaa052
RUN apk --no-cache add ca-certificates

ADD build/esattory_unix /

ENTRYPOINT ["/esattory_unix"]
