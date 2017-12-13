package handler

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/iotdog/microsnail/dbsvc/config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// MgoCli is the default session for mongodb.
	MgoCli *mgo.Session
	// ErrRoleInvalid is the error for user role unauthorized.
	// ErrRoleInvalid = errors.New("user role unauthorized")
)

// InitMongoDB initializes the MongoDB session. 需要重启服务器才能生效
func InitMongoDB() error {
	mgoURL := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		config.Instance().MongoDBAdmin, config.Instance().MongoDBAdmPasswd,
		config.Instance().MongoDBHost, config.Instance().MongoDBPort)

	dialInfo, err := mgo.ParseURL(mgoURL)
	if err != nil {
		return err
	}
	if config.Instance().Mode != config.Debug {
		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = false
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	MgoCli, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}

	return nil
}

type MgoDocument struct {
	ID        bson.ObjectId     `bson:"_id"`
	CreatedAt time.Time         `bson:"createdAt"`
	UpdatedAt time.Time         `bson:"updatedAt"`
	Content   map[string]string `bson:"content"`
}

func NewMgoDocument(content *map[string]string) *MgoDocument {
	return &MgoDocument{
		ID:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Content:   *content,
	}
}

func (md *MgoDocument) Find(database, collection string, query bson.M) error {
	return MgoCli.DB(database).C(collection).Find(query).One(md)
}

func (md *MgoDocument) Upsert(database, collection string) error {
	_, err := MgoCli.DB(database).C(collection).Upsert(bson.M{"_id": md.ID}, md)
	return err
}

func (md *MgoDocument) Delete(database, collection string) error {
	return MgoCli.DB(database).C(collection).Remove(bson.M{"_id": md.ID})
}
