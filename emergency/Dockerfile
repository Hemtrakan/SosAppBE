FROM golang:1.19-alpine AS build_stage

ENV ENV="dev"

RUN apk update && apk upgrade && apk add --no-cache build-base bash git openssh
LABEL maintainer="Chirapon Hemtrakan <chirapon.job@gmail.com>"

ARG PACKAGE_PATH=emergency
RUN mkdir -p /go/src/
WORKDIR /go/src/${PACKAGE_PATH}

COPY . /go/src/${PACKAGE_PATH}/
RUN go mod download

RUN go build -o emergency

ENTRYPOINT ./emergency

EXPOSE 81

CMD ["sh", "-c", "/emergency"]