FROM golang:1.19-alpine AS build_stage

RUN apk update && apk upgrade && apk add --no-cache build-base bash git openssh
LABEL maintainer="Chirapon Hemtrakan <chirapon.job@gmail.com>"

ENV PACKAGE_PATH=hotline
RUN mkdir -p /go/src/
WORKDIR /go/src/$PACKAGE_PATH

COPY . /go/src/$PACKAGE_PATH/
RUN go mod download

RUN go build -o hotline

ENV ENV="dev"

ENTRYPOINT ./hotline
EXPOSE 82
CMD ["sh", "-c", "/hotline"]