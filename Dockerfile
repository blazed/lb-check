FROM alpine:latest

ADD lb /go/bin/lb
ENTRYPOINT /go/bin/lb
