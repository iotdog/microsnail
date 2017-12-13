package main

import (
	"github.com/iotdog/microsnail/dbsvc/config"
	"github.com/iotdog/microsnail/dbsvc/handler"
	dbsvc "github.com/iotdog/microsnail/dbsvc/proto"
	"github.com/iotdog/microsnail/dbsvc/wrapper"
	"github.com/leesper/holmes"
	micro "github.com/micro/go-micro"
)

func main() {
	defer holmes.Start().Stop()

	conf, err := config.LoadConfigs(config.Debug, "./config")
	if err != nil {
		holmes.Errorln(err)
		return
	}

	err = handler.InitMongoDB()
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
