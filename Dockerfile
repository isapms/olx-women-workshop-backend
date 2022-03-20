# syntax=docker/dockerfile:1
FROM golang:1.17.8-alpine

WORKDIR /olx-women-workshop-2022-backend

ADD . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

COPY *.go ./
RUN go build -o /olx-women-workshop-2022-backend

EXPOSE 4040

ENTRYPOINT CompileDaemon -command="./olx-women-workshop-2022-backend"