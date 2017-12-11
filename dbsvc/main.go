package main

import (
	"github.com/iotdog/microsnail/dbsvc/configs"
	"github.com/iotdog/microsnail/dbsvc/handler"
	dbsvc "github.com/iotdog/microsnail/dbsvc/proto"
	"github.com/iotdog/microsnail/dbsvc/wrapper"
	"github.com/leesper/holmes"
	micro "github.com/micro/go-micro"
)

func main() {
	defer holmes.Start().Stop()

	conf, err := configs.LoadConfigs(configs.Debug, "./configs")
	if err != nil {
		holmes.Errorln(err)
		return
	}

	service := micro.NewService(
		micro.Name(conf.ServiceName),
		micro.WrapHandler(wrapper.LogWrapper),
	)
	service.Init()
	dbsvc.RegisterDBServiceHandler(service.Server(), new(handler.DBService))
	if err := service.Run(); err != nil {
		holmes.Errorln(err)
	}
}
