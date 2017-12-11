package main

import (
	"context"

	"github.com/iotdog/microsnail/dbsvc/configs"
	dbsvc "github.com/iotdog/microsnail/dbsvc/proto"
	"github.com/leesper/holmes"
	"github.com/micro/go-micro/client"
)

func main() {
	defer holmes.Start().Stop()

	conf, err := configs.LoadConfigs(configs.Debug, "../configs")
	if err != nil {
		holmes.Errorln(err)
		return
	}

	client := dbsvc.NewDBServiceClient(conf.ServiceName, client.DefaultClient)
	rsp, err := client.Create(context.TODO(), &dbsvc.CreateReq{
		Db:  &dbsvc.Database{},
		Rec: &dbsvc.Record{},
	})
	if err != nil {
		holmes.Errorln(err)
		return
	}
	holmes.Infoln("\ncode: ", rsp.Common.Code, "\nmsg: ", rsp.Common.Message)
}
