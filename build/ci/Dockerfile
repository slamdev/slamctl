FROM alpine

RUN apk add --no-cache git make musl-dev go \
 && apk --no-cache add --repository http://dl-cdn.alpinelinux.org/alpine/edge/main nodejs npm

ENV PATH /root/go/bin:$PATH

RUN go version
RUN node --version
RUN npm --version

WORKDIR /opt
