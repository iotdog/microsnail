package wrapper

import (
	"context"

	"github.com/leesper/holmes"
	"github.com/micro/go-micro/server"
)

func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		holmes.Infof("[%s]: %s called", req.Service(), req.Method())
		return fn(ctx, req, rsp)
	}
}
