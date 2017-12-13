package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	restful "github.com/emicklei/go-restful"
	usersvc "github.com/iotdog/microsnail/usersvc/proto"
	"github.com/micro/go-micro/client"
	web "github.com/micro/go-web"
)

type UserRegParam struct {
	Phone  string `json:"phone"`
	Passwd string `json:"password"`
	Verif  string `json:"verif"`
}

type RestAPI struct{}

var (
	cl usersvc.UserServiceClient
)

func (ra *RestAPI) Register(req *restful.Request, rsp *restful.Response) {
	var input UserRegParam
	err := json.NewDecoder(req.Request.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)

	resp, err := cl.Register(context.TODO(), &usersvc.RegReq{
		Phone:        input.Phone,
		Password:     input.Passwd,
		Verification: input.Verif,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(resp)
	// rsp.WriteJson(map[string]string{
	// 	"test": "have fun",
	// }, "application/json")
}

func main() {

	cl = usersvc.NewUserServiceClient("iotdog.microsnail.usersvc.debug", client.DefaultClient)

	service := web.NewService(
		web.Name("iotdog.microsnail.api.test"),
	)
	service.Init()
	api := new(RestAPI)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path("/api")
	ws.Route(ws.POST("/register").To(api.Register))
	wc.Add(ws)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
