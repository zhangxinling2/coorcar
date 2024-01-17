package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/wechat"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := NewZapLog()
	if err != nil {
		log.Fatal("can not create zap log", logger)
	}
	n, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}
	c := context.Background()
	mo, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect db", zap.Error(err))
	}
	m := dao.NewMongo(mo.Database("coolcar"))
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIdResolve: &wechat.Service{
			AppId:     "wx9fc227d95b260fb3",
			AppSecret: "602e30ce220a6732503f85eea2807f7a",
		},
		Mongo:  m,
		Logger: logger,
	})
	err = s.Serve(n)
	if err != nil {
		logger.Fatal("cannot server", zap.Error(err))
	}
}
func NewZapLog() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()

}
