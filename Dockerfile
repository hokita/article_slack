FROM golang:1.15.6-alpine
# For Raspberry Pi
# FROM arm32v6/golang:1.15.6-alpine

ENV GOBIN /go/bin
ENV GO111MODULE=on
ENV GOPATH=

WORKDIR /go
ADD . /go

RUN go mod tidy
RUN go build -o article_slack *.go

CMD ./article_slack
