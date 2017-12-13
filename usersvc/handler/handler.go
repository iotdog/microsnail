package handler

import (
	"context"
	"strings"
	"time"

	dbsvc "github.com/iotdog/microsnail/dbsvc/proto"
	"github.com/iotdog/microsnail/usersvc/config"
	usersvc "github.com/iotdog/microsnail/usersvc/proto"
)

type UserService struct{}

func (h *UserService) Register(ctx context.Context, in *usersvc.RegReq, out *usersvc.RegResp) error {
	var code int32
	var msg string

	if in.Phone == "" || in.Password == "" || in.Verification == "" ||
		in.Verification != "111111" {
		code = 1
		msg = "invalid parameter"
	} else {
		dbsvcCli := NewDBServiceClient()
		dbname := strings.Replace(config.Instance().ServiceName, ".", "_", -1)
		rsp, err := dbsvcCli.Create(context.TODO(), &dbsvc.CreateReq{
			Db: &dbsvc.Database{
				Name:       dbname,
				Collection: "users",
			},
			Rec: &dbsvc.Record{
				Time: time.Now().Unix(),
				Metadata: map[string]string{
					"phone":    in.Phone,
					"password": in.Password,
				},
			},
		})
		if err != nil {
			code = 1
			msg = err.Error()
		} else {
			code = rsp.Common.Code
			msg = rsp.Common.Message
		}
	}

	out.Common = &usersvc.CommonResp{
		Code:    int32(code),
		Message: msg,
	}

	return nil
}
