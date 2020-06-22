FROM golang:1.12-alpine3.9

WORKDIR /go/src/scaffold
COPY . .

ARG VERSION=1.0.0

ENV GO111MODULE=on

RUN apk update \
    && apk add --update --no-cache --virtual .build-deps \
        python3 python3-dev build-base alpine-sdk freetds-dev musl-dev libffi-dev bash ca-certificates \
        g++ gcc git curl bzr libressl libressl-dev linux-headers make zlib-dev libxslt-dev tzdata supervisor rsyslog \
    && wget https://github.com/edenhill/librdkafka/archive/v$VERSION.tar.gz -O /tmp/librdkafka-$VERSION.tar.gz \
    && pwd \
    && cd /tmp/ \
    && tar zxf librdkafka-$VERSION.tar.gz \
    && cd librdkafka-$VERSION  \
    && ./configure \
    && make \
    && make install \
    && cd /go/src/scaffold \
    && go mod download \
    && rm -r /usr/lib/python*/ensurepip && pip3 install --upgrade pip setuptools  \
    && if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi \
    && ln -sf /usr/bin/python3 /usr/bin/python \
    && rm -r /root/.cache \
    && GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o scaffold-linux

RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

# ADD supervisord.conf /etc/supervisord.conf

EXPOSE 5555

# ENTRYPOINT ["sh", "-c", "supervisord -nc /etc/supervisord.conf"]
