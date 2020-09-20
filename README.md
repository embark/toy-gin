# toy-gin
A toy API for creating book reservations

## Installation
1. You need [Go](https://golang.org/) installed (**version 1.14+ is required**).

2. Build and run the server, dependencies should be automatically installed on the first run. The server will run on http://localhost:8080

```sh
$ go build
$ ./toy-gin
```

## Examples
```sh
$ curl -X POST http://localhost:8080/request -F "title=ABC" -F "email=test"
$ curl -X POST http://localhost:8080/request -F "title=XYZ" -F "email=test"

# TODOs: Will return NotImplemented
$ curl -X GET http://localhost:8080/request -v
$ curl -X GET http://localhost:8080/request/1 -v
$ curl -X DELETE http://localhost:8080/1 -v
```
