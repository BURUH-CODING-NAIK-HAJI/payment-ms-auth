FROM golang

WORKDIR /app

COPY . /app

RUN go mod download

RUN go install github.com/cosmtrek/air@latest
