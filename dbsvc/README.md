# Database Service

service for MongoDB and Couchbase Server.

# Code Framework

## proto

definitions of service protocols using protobuf

### Generate Service Proto

```
protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src \
        $GOPATH/src/github.com/iotdog/microsnail/dbsvc/proto/svc.proto
```

## handler

implementations of service handle functions

## wrapper

service middleware

## client

service client used for testing
