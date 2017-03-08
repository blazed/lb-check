FROM alpine:latest

ADD lb /lb
CMD ["/lb"]
