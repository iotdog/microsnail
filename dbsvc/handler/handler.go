package handler

import (
	"context"

	"gopkg.in/mgo.v2/bson"

	dbsvc "github.com/iotdog/microsnail/dbsvc/proto"
	"github.com/leesper/holmes"
)

type DBService struct{}

func (h *DBService) Create(ctx context.Context, in *dbsvc.CreateReq, out *dbsvc.CreateResp) error {
	code := 0
	msg := "ok"
	if len(in.Rec.Metadata) == 0 {
		code = 1
		msg = "invalid content"
	} else {
		mgoDoc := NewMgoDocument(&in.Rec.Metadata)
		err := mgoDoc.Upsert(in.Db.Name, in.Db.Collection)
		if err != nil {
			holmes.Errorln(err)
			code = 1
			msg = err.Error()
		}
	}
	out.Common = &dbsvc.CommonResp{
		Code:    int32(code),
		Message: msg,
	}

	return nil
}

func (h *DBService) Delete(ctx context.Context, in *dbsvc.DeleteReq, out *dbsvc.DeleteResp) error {
	code := 0
	msg := "ok"
	if !bson.IsObjectIdHex(in.Recid) {
		code = 1
		msg = "invalid document id"
	} else {
		mgoDoc := new(MgoDocument)
		err := mgoDoc.Find(in.Db.Name, in.Db.Collection, bson.M{"_id": bson.ObjectIdHex(in.Recid)})
		if err != nil {
			code = 1
			msg = err.Error()
		} else {
			err = mgoDoc.Delete(in.Db.Name, in.Db.Collection)
			if err != nil {
				code = 1
				msg = err.Error()
			}
		}
	}
	out.Common = &dbsvc.CommonResp{
		Code:    int32(code),
		Message: msg,
	}
	return nil
}
