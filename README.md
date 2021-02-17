# goSimpleFTP

## Requirements

https://github.com/goftp/server

Attacker:
* Golang

Victim:
* Shell access

## Installation

### Build for Windows machine:
* `$ go get github.com/goftp/file-driver`
* `$ go get github.com/goftp/server`
* `$ GOOS=windows GOARCH=amd64 go build main.go` 64bit
* `$ GOOS=windows GOARCH=386 go build main.go` 32bit

### Build for Linux machine:
* `$ go get github.com/goftp/file-driver`
* `$ go get github.com/goftp/server`
* `$ GOOS=linux GOARCH=amd64 go build main.go` 64bit
* `$ GOOS=linux GOARCH=386 go build main.go` 32bit
