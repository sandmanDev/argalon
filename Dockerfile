FROM golang:1.16.4

ENV GIN_MODE=release
ENV PORT=8088

ADD . /go/github.com/argalon

WORKDIR /go/github.com/argalon
RUN go mod tidy

RUN go build

EXPOSE $PORT

ENTRYPOINT ./argalon