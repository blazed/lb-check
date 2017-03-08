FROM alpine:latest

ADD lb /lb
ENTRYPOINT /lb
