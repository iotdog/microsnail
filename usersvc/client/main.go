package main

import (
	"context"

	"github.com/iotdog/microsnail/usersvc/config"
	usersvc "github.com/iotdog/microsnail/usersvc/proto"
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

	client := usersvc.NewUserServiceClient(conf.ServiceName, client.DefaultClient)
	// user registration
	rsp, err := client.Register(context.TODO(), &usersvc.RegReq{
		Phone:        "18084454865",
		Password:     "123456",
		Verification: "111111",
	})

	if err != nil {
		holmes.Errorln(err)
		return
	}
	holmes.Infoln("\ncode: ", rsp.Common.Code, "\nmsg: ", rsp.Common.Message)
}
