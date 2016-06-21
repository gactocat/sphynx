# snowshoe

A boilerplate for Go + Echo + Godep.

## Get Start

```
$ go get -u github.com/tools/godep
$ go get -d github.com/gactocat/snowshoe
$ cd ${GOPATH}/src/github.com/gactocat/snowshoe
```

## APIs

```
$ curl -v -X POST -H'Content-Type: application/json' -d'{"name":"Paul Young"}' localhost:1323/users
$ curl -v localhost:1323/users/1
$ curl -v -X PUT -H'Content-Type: application/json' -d'{"name":"Paul Smith"}' localhost:1323/users/1
$ curl -v -X DELETE localhost:1323/users/1
```


