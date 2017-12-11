package handler

import (
	"context"

	dbsvc "github.com/iotdog/microsnail/dbsvc/proto"
)

type DBService struct{}

func (h *DBService) Create(ctx context.Context, in *dbsvc.CreateReq, out *dbsvc.CreateResp) error {
	out.Common = &dbsvc.CommonResp{
		Code:    0,
		Message: "Have Fun",
	}
	return nil
}
