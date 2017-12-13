package handler

import (
	dbsvc "github.com/iotdog/microsnail/dbsvc/proto"
	"github.com/iotdog/microsnail/usersvc/config"
	"github.com/micro/go-micro/client"
)

func NewDBServiceClient() dbsvc.DBServiceClient {
	return dbsvc.NewDBServiceClient(config.Instance().DBServiceName, client.DefaultClient)
}
