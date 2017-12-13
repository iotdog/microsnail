package main

import (
	"github.com/iotdog/microsnail/usersvc/config"
	"github.com/iotdog/microsnail/usersvc/handler"
	usersvc "github.com/iotdog/microsnail/usersvc/proto"
	"github.com/iotdog/microsnail/usersvc/wrapper"
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

	service := micro.NewService(
		micro.Name(conf.ServiceName),
		micro.WrapHandler(wrapper.LogWrapper),
	)
	service.Init()
	usersvc.RegisterUserServiceHandler(service.Server(), new(handler.UserService))
	if err := service.Run(); err != nil {
		holmes.Errorln(err)
	}
}
