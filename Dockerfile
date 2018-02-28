FROM alpine:3.5

RUN mkdir /app
WORKDIR /app

ADD ./goreplay-installer /usr/bin

ENTRYPOINT [ "/usr/bin/goreplay-installer"]
