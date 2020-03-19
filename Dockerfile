FROM golang:alpine

RUN apk update && apk add git curl

RUN go get github.com/oxequa/realize

WORKDIR /opt/app

COPY go.* /opt/app/

RUN go mod download
RUN go mod vendor

COPY . .

EXPOSE 8080