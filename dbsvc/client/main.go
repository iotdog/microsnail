package main

import (
	"context"

	"github.com/iotdog/microsnail/dbsvc/config"
	dbsvc "github.com/iotdog/microsnail/dbsvc/proto"
	"github.com/leesper/holmes"
	"github.com/micro/go-micro/client"
)

func main() {
	defer holmes.Start().Stop()

	conf, err := config.LoadConfigs(config.Debug, "../config")
	if err != nil {
		holmes.Errorln(err)
		return
	}

	client := dbsvc.NewDBServiceClient(conf.ServiceName, client.DefaultClient)
	// create new document
	// rsp, err := client.Create(context.TODO(), &dbsvc.CreateReq{
	// 	Db: &dbsvc.Database{
	// 		Name:       "test",
	// 		Collection: "a",
	// 	},
	// 	Rec: &dbsvc.Record{
	// 		Time: time.Now().Unix(),
	// 		Metadata: map[string]string{
	// 			"hello": "world",
	// 		},
	// 	},
	// })
	// delete document
	rsp, err := client.Delete(context.TODO(), &dbsvc.DeleteReq{
		Db: &dbsvc.Database{
			Name:       "test",
			Collection: "a",
		},
		Recid: "5a30c72565598b1fa0052a5a",
	})

	if err != nil {
		holmes.Errorln(err)
		return
	}
	holmes.Infoln("\ncode: ", rsp.Common.Code, "\nmsg: ", rsp.Common.Message)
}
