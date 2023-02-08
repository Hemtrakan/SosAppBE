FROM golang:1.18-alpine AS build_stage

RUN apk update && apk upgrade && apk add --no-cache build-base bash git openssh
LABEL maintainer="Chirapon Hemtrakan <chirapon.job@gmail.com>"

ENV PACKAGE_PATH=SosApp
RUN mkdir -p /go/src/
WORKDIR /go/src/$PACKAGE_PATH

COPY . /go/src/$PACKAGE_PATH/
RUN go mod download

RUN go build -o bin/SosApp

ENV ENV="dev"

ENTRYPOINT ./bin/SosApp
EXPOSE 80
CMD ["sh", "-c", "/bin/SosApp"]