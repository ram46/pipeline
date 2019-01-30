FROM golang:latest

WORKDIR /go/src/app

COPY helloworld . 

CMD ["app"]
