FROM golang:latest

COPY . /go/src/github.com/mickelsonm/demo-api

WORKDIR /go/src/github.com/mickelsonm/demo-api

RUN go get && go build -o API ./main.go

ENTRYPOINT /go/src/github.com/mickelsonm/demo-api/API

EXPOSE 8080
