# usersvc Service

Service introduction

# Code Framework

## proto

definitions of service protocols using protobuf

### Generate Service Proto

```
protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src \
        $GOPATH/src/github.com/iotdog/microsnail/usersvc/proto/svc.proto
```

## handler

implementations of service handle functions

## wrapper

service middleware

## configs

service configuration

## client

service client used for testing

# Run Code

* start consul

```
consul agent -dev
```

* run server

```
cd $GOPATH/src/github.com/iotdog/microsnail/usersvc
go run main.go
```

* run client

```
cd $GOPATH/src/github.com/iotdog/microsnail/usersvc/client
go run main.go
```
