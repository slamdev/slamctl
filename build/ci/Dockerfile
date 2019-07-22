FROM alpine

RUN apk add --no-cache git make musl-dev go bash curl docker \
 && apk --no-cache add --repository http://dl-cdn.alpinelinux.org/alpine/edge/main nodejs npm \
 && go version && node --version && npm --version

ENV PATH /root/go/bin:$PATH

WORKDIR /opt
