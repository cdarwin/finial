FROM library/golang:1.6.0-alpine

RUN apk add --update git \
    && rm -rf /var/cache/apk/*
