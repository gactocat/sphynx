FROM golang:1.6

MAINTAINER gactocat <gactocat@gmail.com>

RUN go get -u gopkg.in/godo.v2/cmd/godo

WORKDIR /go/src/github.com/gactocat/snowshoe2

EXPOSE 8080

CMD ["go", "run", "main.go"]

