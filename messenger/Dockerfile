FROM golang:1.19-alpine AS build_stage

ENV ENV="dev"

#"DEVELOP"
#"DEVSPACE"
#"PRODUCTION"
ENV FLAVOR="DEVELOP"

RUN apk update && apk upgrade && apk add --no-cache build-base bash git openssh
LABEL maintainer="Chirapon Hemtrakan <chirapon.job@gmail.com>"

ARG PACKAGE_PATH=messenger
RUN mkdir -p /go/src/
WORKDIR /go/src/${PACKAGE_PATH}

COPY . /go/src/${PACKAGE_PATH}/
RUN go mod download

RUN go build -o messenger

ENTRYPOINT ./messenger

EXPOSE 83

CMD ["sh", "-c", "/messenger"]