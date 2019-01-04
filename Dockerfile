FROM alpine:3.5
MAINTAINER nagaa052

ADD build/esattory_unix /

ENTRYPOINT ["/esattory_unix"]
