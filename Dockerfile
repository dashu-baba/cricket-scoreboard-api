FROM golang:alpine

RUN apk update && apk add git curl

RUN go get github.com/oxequa/realize
RUN go get -u github.com/swaggo/swag/cmd/swag

WORKDIR /opt/app

COPY go.* /opt/app/

RUN go mod download
RUN go mod vendor

COPY . .

RUN swag init --dir ./src --output ./src/docs

EXPOSE 8080